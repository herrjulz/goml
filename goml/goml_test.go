package goml_test

import (
	"fmt"

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

			Set(yml, "mapArray.foo:bar.arr.0", "new")
			Expect(Get(yml, "mapArray.foo:bar.arr.0")).To(Equal("new"))
		})
	})

	Context("Delete", func() {
		It("should delete a value from a map, array, array with maps", func() {
			yml, _ := simpleyaml.NewYaml([]byte(yaml))
			yml, err := Delete(yml, "map.name")
			if err != nil {
				fmt.Println(err)
			}

		})
	})

})
