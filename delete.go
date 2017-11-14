package goml

import (
	"errors"
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/smallfish/simpleyaml"
)

func DeleteInFile(file string, path string) error {
	yaml, err := ReadYamlFromFile(file)
	if err != nil {
		return err
	}

	err = Delete(yaml, path)
	if err != nil {
		return errors.New("Couldn't delete property for path: " + path)
	}

	if yaml.IsMap() {
		keys, err := yaml.GetMapKeys()
		if err != nil {
			return err
		}

		if len(keys) == 0 {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}
	}

	return WriteYaml(yaml, file)
}

func DeleteInMemory(file []byte, path string) ([]byte, error) {
	yml, err := simpleyaml.NewYaml(file)
	if err != nil {
		return nil, err
	}

	err = Delete(yml, path)
	if err != nil {
		return nil, errors.New("Couldn't delete property for path: " + path)
	}

	ymlMap, err := yml.Map()
	if err != nil {
		return nil, err
	}

	ymlBytes, err := yaml.Marshal(ymlMap)
	if err != nil {
		return nil, err
	}

	return ymlBytes, nil
}

func Delete(yml *simpleyaml.Yaml, path string) error {
	propsArr := strings.Split(path, ".")
	propName := propsArr[len(propsArr)-1]
	props := propsArr[:len(propsArr)-1]
	newPath := strings.Join(props, ".")

	if index, err := strconv.Atoi(propName); err == nil {
		tmp, props := get(yml, newPath)
		prop, err := tmp.Array()
		if err != nil {
			return err
		}

		prop = append(prop[:index], prop[index+1:]...)

		updateYaml(yml, props, prop)
		return nil
	}

	if strings.Contains(propName, ":") {
		tmp, props := get(yml, newPath)
		prop, err := tmp.Array()
		if err != nil {
			return err
		}

		index, err := returnIndexForProp(propName, prop)
		if err != nil {
			return err
		}

		prop = append(prop[:index], prop[index+1:]...)
		updateYaml(yml, props, prop)
		return nil
	}

	var res map[interface{}]interface{}
	if len(propsArr) == 1 {
		tmp, err := yml.Map()
		if err != nil {
			return err
		}
		res = tmp
	} else {
		prop, _ := get(yml, newPath)
		tmp, err := prop.Map()
		if err != nil {
			return err
		}
		res = tmp
	}
	_, ok := res[propName]
	if ok {
		delete(res, propName)
	} else {
		return errors.New("property not found")
	}

	return nil
}
