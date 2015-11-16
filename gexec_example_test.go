package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("exitCode flag", func() {
	var pathToExecutable string
	var command *exec.Cmd

	BeforeSuite(func() {
		var err error
		pathToExecutable, err = gexec.Build("github.com/amitkgupta/gexec-example")
		Ω(err).ShouldNot(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	Context("when exitCode flag is passed explicitly", func() {
		BeforeEach(func() {
			command = exec.Command(pathToExecutable, "-exitCode=33")
		})

		It("exits with the passed in value", func() {
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(33))
		})
	})

	Context("When exitCode flag is not passed", func() {
		BeforeEach(func() {
			command = exec.Command(pathToExecutable)
		})

		It("exits with the default value", func() {
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(session).Should(gexec.Exit(0))
		})
	})
})
