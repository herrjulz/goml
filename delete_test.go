package goml_test

import (
	. "github.com/JulzDiverse/goml"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/smallfish/simpleyaml"
)

var _ = Describe("Delete", func() {
	var yml *simpleyaml.Yaml
	var err error

	BeforeEach(func() {
		yaml := `map:
  name: foo

array:
- bar
- var
- zar

mapArray:
- foo: bar
  zoo: lion
  arr:
  - one
  - two
  - three
- foo: var
  boo: laa`

		yml, err = simpleyaml.NewYaml([]byte(yaml))
		Expect(err).NotTo(	HaveOccurred())
	})

	It("should delete a value from a map", func() {
		err = Delete(yml, "map.name")
		Expect(err).NotTo(HaveOccurred())

		_, err = Get(yml, "map.name")
		Expect(err).To(MatchError("property not found"))
	})

	It("should delete a value from an array based on name", func() {
		err = Delete(yml, "array.bar")
		Expect(err).NotTo(HaveOccurred())
	})

	It("should delete a value from an array based on index", func() {
		err = Delete(yml, "array.0")
		Expect(err).NotTo(HaveOccurred())

		_, err := Get(yml, "array.:bar")
		Expect(err).To(MatchError("property not found"))

		err = Delete(yml, "array.:zar")
		Expect(err).NotTo(HaveOccurred())

		_, err = Get(yml, "array.:zar")
		Expect(err).To(MatchError("property not found"))
	})

	It("should delete a value from an map inside an array ", func() {
		err = Delete(yml, "mapArray.foo:bar.zoo")
		Expect(err).NotTo(HaveOccurred())

		_, err := Get(yml, "mapArray.foo:bar.zoo")
		Expect(err).To(HaveOccurred())
	})

	It("should delete a value from an array inside a map which in turn is inside an array ", func() {
		err = Delete(yml, "mapArray.foo:bar.arr.0")
		Expect(err).NotTo(HaveOccurred())

		_, err := Get(yml, "mapArray.foo:bar.arr.:one")
		Expect(err).To(MatchError("property not found"))

		err = Delete(yml, "mapArray.foo:bar.arr.:two")
		Expect(err).NotTo(HaveOccurred())

		_, err = Get(yml, "mapArray.foo:bar.arr.:two")
		Expect(err).To(MatchError("property not found"))
	})
})
