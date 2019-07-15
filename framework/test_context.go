package framework

const (
    defaultHost = "http://127.0.0.1:8080"

    // DefaultNumNodes is the number of nodes. If not specified, then number of nodes is auto-detected.
    DefaultNumNodes = -1
)

type TestContextType struct {
    KubeConfig string
    KubeContext string
    KubeAPIContentType string
    KubeVolumeDir string
    CertDir string
    Host string

}

var TestContext TestContextType