package goml_test

import (
	. "github.com/JulzDiverse/goml"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/smallfish/simpleyaml"
)

var _ = Describe("Get", func() {
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
		Expect(err).NotTo(HaveOccurred())
	})

	It("should get a value from a map, array, array with maps", func() {
		value, err := Get(yml, "map.name")
		Expect(err).NotTo(HaveOccurred())
		Expect(value).To(Equal("foo"))

		value, err = Get(yml, "array.0")
		Expect(err).NotTo(HaveOccurred())
		Expect(value).To(Equal("bar"))

		value, err = Get(yml, "mapArray.0.foo")
		Expect(err).NotTo(HaveOccurred())
		Expect(value).To(Equal("bar"))

		value, err = Get(yml, "mapArray.foo:var.boo")
		Expect(err).NotTo(HaveOccurred())
		Expect(value).To(Equal("laa"))

		value, err = Get(yml, "array.:var")
		Expect(err).NotTo(HaveOccurred())
		Expect(value).To(Equal("var"))

		value, err = Get(yml, "mapArray.foo:bar.arr.0")
		Expect(err).NotTo(HaveOccurred())
		Expect(value).To(Equal("one"))
	})
})
