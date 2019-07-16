module github.com/zryfish/framework

go 1.12

replace (
	k8s.io/api => github.com/kubernetes/api v0.0.0-20190620084959-7cf5895f2711
	k8s.io/apiextensions-apiserver => github.com/kubernetes/apiextensions-apiserver v0.0.0-20190620085554-14e95df34f1f
	k8s.io/apimachinery => github.com/kubernetes/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/apiserver => github.com/kubernetes/apiserver v0.0.0-20190620085212-47dc9a115b18
	k8s.io/cli-runtime => github.com/kubernetes/cli-runtime v0.0.0-20190620085706-2090e6d8f84c
	k8s.io/client-go => github.com/kubernetes/client-go v0.0.0-20190620085101-78d2af792bab
	k8s.io/cloud-provider => github.com/kubernetes/cloud-provider v0.0.0-20190620090043-8301c0bda1f0
	k8s.io/cluster-bootstrap => github.com/kubernetes/cluster-bootstrap v0.0.0-20190620090013-c9a0fc045dc1
	k8s.io/code-generator => github.com/kubernetes/code-generator v0.0.0-20190612205613-18da4a14b22b
	k8s.io/component-base => github.com/kubernetes/component-base v0.0.0-20190620085130-185d68e6e6ea
	k8s.io/cri-api => github.com/kubernetes/cri-api v0.0.0-20190531030430-6117653b35f1
	k8s.io/csi-translation-lib => github.com/kubernetes/csi-translation-lib v0.0.0-20190620090116-299a7b270edc
	k8s.io/kube-aggregator => github.com/kubernetes/kube-aggregator v0.0.0-20190620085325-f29e2b4a4f84
	k8s.io/kube-controller-manager => github.com/kubernetes/kube-controller-manager v0.0.0-20190620085942-b7f18460b210
	k8s.io/kube-proxy => github.com/kubernetes/kube-proxy v0.0.0-20190620085809-589f994ddf7f
	k8s.io/kube-scheduler => github.com/kubernetes/kube-scheduler v0.0.0-20190620085912-4acac5405ec6
	k8s.io/kubelet => github.com/kubernetes/kubelet v0.0.0-20190620085838-f1cb295a73c9
	k8s.io/kubernetes => github.com/kubernetes/kubernetes v1.15.0
	k8s.io/legacy-cloud-providers => github.com/kubernetes/legacy-cloud-providers v0.0.0-20190620090156-2138f2c9de18
	k8s.io/metrics => github.com/kubernetes/metrics v0.0.0-20190620085625-3b22d835f165
	k8s.io/sample-apiserver => github.com/kubernetes/sample-apiserver v0.0.0-20190620085408-1aef9010884e
)

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/kubernetes v0.0.0-00010101000000-000000000000
)
