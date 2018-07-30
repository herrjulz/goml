package goml_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/goml"
)

var _ = FDescribe("Paths", func() {
	Context("When retrieving paths of a yaml file", func() {

		var (
			yaml string
			err  error
		)

		//JustBeforeEach(func() {
		//sl, _ := GetPaths([]byte(yaml))
		//paths = sl
		//sl = []string{}
		//fmt.Println("PATHS for test", test, paths)
		//})

		//AfterEach(func() {
		//paths = nil
		//})

		Context("For a simple map", func() {
			BeforeEach(func() {
				yaml = `---
map:
  name: julz`

			})

			It("should not fail", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return the paths", func() {
				paths, err := GetPaths([]byte(yaml))
				Expect(err).ToNot(HaveOccurred())
				Expect(paths[0]).To(Equal("map.name"))
			})

			It("should be a valid goml path", func() {
				paths, err := GetPaths([]byte(yaml))
				Expect(err).ToNot(HaveOccurred())
				result, err := GetInMemory([]byte(yaml), paths[0])
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal("julz"))
			})
		})

		Context("For a simple array", func() {
			BeforeEach(func() {
				yaml = `---
array:
- julz`
			})

			It("should not fail", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return the path", func() {
				paths, err := GetPaths([]byte(yaml))
				Expect(err).ToNot(HaveOccurred())
				fmt.Println("PATHS", paths)
				Expect(paths[0]).To(Equal("array.0"))
			})

			It("should be a valid goml path", func() {
				paths, err := GetPaths([]byte(yaml))
				Expect(err).ToNot(HaveOccurred())
				result, err := GetInMemory([]byte(yaml), paths[0])
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal("julz"))
			})
		})
	})
})
