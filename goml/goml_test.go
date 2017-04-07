package goml_test

import (
	. "github.com/JulzDiverse/goml/goml"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/smallfish/simpleyaml"
)

var _ = Describe("Goml", func() {

	var yaml string
	BeforeEach(func() {
		yaml = `map:
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
	})

	Context("Get", func() {
		It("should get a value from a map, array, array with maps", func() {
			yml, _ := simpleyaml.NewYaml([]byte(yaml))

			Expect(Get(yml, "map.name")).To(Equal("foo"))
			Expect(Get(yml, "array.0")).To(Equal("bar"))
			Expect(Get(yml, "mapArray.0.foo")).To(Equal("bar"))
			Expect(Get(yml, "mapArray.foo:var.boo")).To(Equal("laa"))
			Expect(Get(yml, "array.:var")).To(Equal("var"))
			Expect(Get(yml, "mapArray.foo:bar.arr.0")).To(Equal("one"))

		})
	})

	Context("Set", func() {
		It("should add an value to an array", func() {
			yml, _ := simpleyaml.NewYaml([]byte(yaml))

			Set(yml, "array.2", "bumblebee")
			Set(yml, "array.+", "optimusPrime")
			Set(yml, "mapArray.0.foo", "wolverine")
			Set(yml, "mapArray.foo:var.boo", "baymax")

			Expect(Get(yml, "array.2")).To(Equal("bumblebee"))
			Expect(Get(yml, "array.3")).To(Equal("optimusPrime"))
			Expect(Get(yml, "mapArray.0.foo")).To(Equal("wolverine"))
			Expect(Get(yml, "mapArray.foo:var.boo")).To(Equal("baymax"))

			Set(yml, "array.:optimusPrime", "pikachu")
			Expect(Get(yml, "array.:pikachu")).To(Equal("pikachu"))
			Set(yml, "mapArray.foo:wolverine.arr.0", "new")
			Expect(Get(yml, "mapArray.foo:wolverine.arr.0")).To(Equal("new"))
		})
	})

	Context("Delete", func() {
		It("should delete a value from a map", func() {
			yml, _ := simpleyaml.NewYaml([]byte(yaml))
			Delete(yml, "map.name")
			_, err := Get(yml, "map.name")
			Expect(err).NotTo(BeNil())
		})

		It("should delete a value from an array ", func() {
			yml, _ := simpleyaml.NewYaml([]byte(yaml))
			Delete(yml, "array.0")
			_, err := Get(yml, "array.:bar")
			Expect(err).NotTo(BeNil())

			Delete(yml, "array.:zar")
			_, err = Get(yml, "array.:zar")
			Expect(err).NotTo(BeNil())

		})

		It("should delete a value from an map inside an array ", func() {
			yml, _ := simpleyaml.NewYaml([]byte(yaml))
			Delete(yml, "array.mapArray.foo:bar.zoo")
			_, err := Get(yml, "array.mapArray.foo:bar.zoo")
			Expect(err).NotTo(BeNil())
		})

		It("should delete a value from an array inside a map which in turn is inside an array ", func() {
			yml, _ := simpleyaml.NewYaml([]byte(yaml))
			Delete(yml, "array.mapArray.foo:bar.arr.0")
			_, err := Get(yml, "array.mapArray.foo:bar.arr.:one")
			Expect(err).NotTo(BeNil())

			Delete(yml, "array.mapArray.foo:bar.arr.:two")
			_, err = Get(yml, "array.mapArray.foo:bar.arr.:two")
			Expect(err).NotTo(BeNil())
		})
	})

})
