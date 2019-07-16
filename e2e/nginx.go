package e2e

import (
    "github.com/onsi/ginkgo"
    "github.com/onsi/gomega"
    "github.com/zryfish/framework/framework"
)

var _ = ginkgo.Describe("nginx service", func() {
    f := framework.NewDefaultFramework("nginx")

    ginkgo.It("should do", func() {
        gomega.Expect(f.Namespace.Name).To(gomega.ContainSubstring("e2e"))
    })

    ginkgo.It("should not do", func() {
        gomega.Expect(f.Namespace.Name).To(gomega.ContainSubstring("e2e"))
    })
})
