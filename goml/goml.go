package goml

import (
	"io/ioutil"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/smallfish/simpleyaml"
)

func Get(yml *simpleyaml.Yaml, path string) (interface{}, error) {
	res, err := get(yml, path).String()
	if err != nil {
		return "", err
	}
	return res, nil
}

func Set(yml *simpleyaml.Yaml, path string, val interface{}) error {
	props := strings.Split(path, ".")
	propName := props[len(props)-1]
	props = props[:len(props)-1]
	newPath := strings.Join(props, ".")

	if index, err := strconv.Atoi(propName); err == nil {
		prop, err := get(yml, newPath).Array()
		if err != nil {
			return err
		}

		prop[index] = val
		updateYaml(*yml, props, prop)
		return nil
	}

	if propName == "+" {
		prop, err := get(yml, newPath).Array()
		if err != nil {
			return err
		}

		prop = append(prop, val)
		updateYaml(*yml, props, prop)
		return nil
	}

	if strings.Contains(propName, ":") {
		prop, err := get(yml, newPath).Array()
		if err != nil {
			return err
		}

		index := returnIndexForProp(propName, prop)
		prop[index] = val
		updateYaml(*yml, props, prop)
		return nil
	}

	prop, err := get(yml, newPath).Map()
	if err != nil {
		return err
	}

	prop[propName] = val

	return nil
}

func Delete(yml *simpleyaml.Yaml, path string) (*simpleyaml.Yaml, error) {
	props := strings.Split(path, ".")
	propName := props[len(props)-1]
	props = props[:len(props)-1]
	newPath := strings.Join(props, ".")

	res, err := get(yml, newPath).Map()
	if err != nil {
		return nil, err
	}

	_, ok := res[propName]
	if ok {
		delete(res, propName)
	}

	return yml, nil
}

func WriteYaml(yml *simpleyaml.Yaml, file string) error {
	goml, err := yml.Map()
	if err != nil {
		return err
	}

	gomlSave, err := yaml.Marshal(goml)
	if err != nil {
		return err
	}

	ioutil.WriteFile(file, gomlSave, 0644)

	return nil
}

func ReadYaml(yaml []byte) (*simpleyaml.Yaml, error) {
	yml, err := simpleyaml.NewYaml(yaml)
	if err != nil {
		return nil, err
	}
	return yml, nil
}

func ReadYamlFromFile(filename string) (*simpleyaml.Yaml, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	yml, _ := simpleyaml.NewYaml(file)
	return yml, nil
}

func get(yml *simpleyaml.Yaml, path string) *simpleyaml.Yaml {
	props := strings.Split(path, ".")
	for _, p := range props {
		if index, err := strconv.Atoi(p); err == nil {
			yml = yml.GetIndex(index)
			continue
		}

		if strings.Contains(p, ":") {
			if prop, err := yml.Array(); err == nil {
				index := returnIndexForProp(p, prop)
				yml = yml.GetIndex(index)
				continue
			}
		}
		yml = yml.Get(p)
	}
	return yml
}

func updateYaml(yml simpleyaml.Yaml, props []string, prop []interface{}) {
	propName := props[len(props)-1]
	props = props[:len(props)-1]
	yaml, _ := yml.GetPath(props...).Map()

	yaml[propName] = prop
}

func returnIndexForProp(propName string, array []interface{}) int {
	var index int

	keyVal := strings.Split(propName, ":")
	key, val := keyVal[0], keyVal[1]

	for i, _ := range array {
		if key == "" {
			check := array[i]
			if check == val {
				index = i
				break
			}
		} else {
			check := array[i].(map[interface{}]interface{})
			if check[key] == val {
				index = i
				break
			}
		}
	}

	return index
}
