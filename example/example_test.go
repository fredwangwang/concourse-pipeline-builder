package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var _ = Describe("Example", func() {
	It("runs", func() {
		pathToMain, err := gexec.Build("github.com/fredwangwang/concourse-pipeline-builder/example")
		Expect(err).NotTo(HaveOccurred())

		cmd := exec.Command(pathToMain)
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))
	})
})
