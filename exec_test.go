package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/tlwr/merge-mgr"
)

var _ = Describe("exec", func() {
	It("returns true for existing executables", func() {
		found := IsExecutable("ls")
		Expect(found).To(Equal(true))
	})

	It("returns false for non-existent executables", func() {
		found := IsExecutable("this-executable-does-not-exist")
		Expect(found).To(Equal(false))
	})
})
