package framework

import (
    "github.com/onsi/ginkgo"
    "github.com/onsi/gomega"
    "k8s.io/apimachinery/pkg/runtime/schema"

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
        ns, err := f.Create
    }
}

func (f *Framework) CreateNamespace(baseName string, labels map[string]string) (*v1.Namespace, error) {
    ns, err :=
}