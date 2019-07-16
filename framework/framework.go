package framework

import (
    "fmt"
    "github.com/onsi/ginkgo"
    "github.com/onsi/gomega"
    "k8s.io/apimachinery/pkg/api/errors"
    "k8s.io/apimachinery/pkg/runtime/schema"
    "strings"

    "k8s.io/api/core/v1"
    "k8s.io/client-go/dynamic"
	clientset "k8s.io/client-go/kubernetes"
)

type Framework struct {
	BaseName string

	// ClientSet uses internal objects, you should use ClientSet where possible.
	ClientSet  clientset.Interface
	DynamicClient dynamic.Interface

	// configuration for framework's client
	Options               Options
	SkipNamespaceCreation bool

	Namespace          *v1.Namespace
	namespacesToDelete []*v1.Namespace
}

type Options struct {
    ClientQPS float32
    ClientBurst int
    GroupVersion *schema.GroupVersion
}

func NewDefaultFramework(baseName string) *Framework {
    options := Options{
        ClientQPS: 20,
        ClientBurst: 50,
    }

    return NewFramework(baseName, options, nil)
}

func NewFramework(baseName string, options Options, client clientset.Interface) *Framework {
    f := &Framework{
        BaseName: baseName,
        ClientSet: client,
        Options: options,
    }

    ginkgo.BeforeEach(f.BeforeEach)
    ginkgo.AfterEach(f.AfterEach)

    return f
}

// BeforeEach gets a clientset and makes a namespace
func (f *Framework) BeforeEach() {
    if f.ClientSet == nil {
        ginkgo.By("Creating a kubernetes client")
        config, err := LoadConfig()
        gomega.Expect(err).NotTo(gomega.HaveOccurred())
        f.ClientSet, err = clientset.NewForConfig(config)
        gomega.Expect(err).NotTo(gomega.HaveOccurred())
    }

    if !f.SkipNamespaceCreation {
        ns, err := f.CreateNamespace(f.BaseName, map[string]string{
            "e2e-framework": f.BaseName,
        })
        gomega.Expect(err).NotTo(gomega.HaveOccurred(), "failed to create namespace")
        ginkgo.By(fmt.Sprintf("Create namespace %s successfully", ns.Name))

        f.Namespace = ns
    }
}

func (f *Framework) AfterEach()  {
    defer func() {
        nsDeletionErrors := map[string]error{}

        if TestContext.DeleteNamespace && (TestContext.DeleteNamespaceOnFailure || !ginkgo.CurrentGinkgoTestDescription().Failed) {
            for _, ns := range f.namespacesToDelete {
                ginkgo.By(fmt.Sprintf("Destroying namespace %q for this suite", ns.Name))
                if err := deleteNS(f.ClientSet, f.DynamicClient, ns.Name, DefaultNamespaceDeletionTimeout); err != nil {
                    if !errors.IsNotFound(err) {
                        nsDeletionErrors[ns.Name] = err
                    } else {
                        Logf("Namespace %v was already deleted", ns.Name)
                    }
                }
            }
        } else {
            if !TestContext.DeleteNamespace {
                Logf("Found DeleteNamespace=false, skipping namespace deletion!")
            } else {
                Logf("Found DeleteNamespaceOnFailure=false and current test failed, skipping namespace deletion!")
            }
        }

        f.Namespace = nil
        f.ClientSet = nil
        f.namespacesToDelete = nil

        if len(nsDeletionErrors) > 0 {
            messages := []string{}
            for namespaceKey, namespaceErr := range nsDeletionErrors {
                messages = append(messages, fmt.Sprintf("Couldn't delete ns: %q: %s (%#v)", namespaceKey, namespaceErr, namespaceErr))
            }
            ginkgo.Fail(strings.Join(messages, ","))
        }
    }()
}

func (f *Framework) CreateNamespace(baseName string, labels map[string]string) (*v1.Namespace, error) {
    ns, err := CreateTestingNS(f.BaseName, f.ClientSet, labels)

    // check ns instead of error or see if its nil as we may
    // fail to create serviceaccount in it.
    // In this case. we should not forget to delete the namespace
    if ns != nil {
        f.namespacesToDelete = append(f.namespacesToDelete, ns)
    }
    return ns, err
}

