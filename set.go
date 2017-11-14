package goml

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/smallfish/simpleyaml"
)

func SetKeyInFile(file string, path string, key string) error {
	yaml, err := ReadYamlFromFile(file)
	if err != nil {
		return err
	}

	err = SetKey(yaml, path, key)
	if err != nil {
		return err
	}

	return WriteYaml(yaml, file)
}

func SetKey(yaml *simpleyaml.Yaml, path string, key string) error {
	bytes, err := ioutil.ReadFile(key)
	if err != nil {
		return err
	}
	return Set(yaml, path, string(bytes))
}

func SetInFile(file string, path string, val interface{}) error {
	yaml, err := ReadYamlFromFile(file)
	if err != nil {
		return err
	}

	err = Set(yaml, path, val)
	if err != nil {
		return err
	}

	return WriteYaml(yaml, file)
}

func SetInMemory(file []byte, path string, val interface{}) ([]byte, error) {
	yml, err := simpleyaml.NewYaml(file)
	if err != nil {
		return nil, err
	}

	err = Set(yml, path, val)
	if err != nil {
		return nil, err
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

func Set(yml *simpleyaml.Yaml, path string, val interface{}) error {
	propsArr := strings.Split(path, ".")
	propName := propsArr[len(propsArr)-1]
	props := propsArr[:len(propsArr)-1]
	newPath := strings.Join(props, ".")

	if index, err := strconv.Atoi(propName); err == nil {
		tmp, props := get(yml, newPath)
		if props == nil {
			return errors.New("property not found")
		}

		prop, err := tmp.Array()
		if err != nil {
			return err
		}

		prop[index] = convertValueType(val)

		updateYaml(yml, props, prop)
		return nil
	}

	if propName == "+" {
		tmp, props := get(yml, newPath)
		prop, err := tmp.Array()
		if err != nil {
			return err
		}
		resolved := convertValueType(val)
		prop = append(prop, resolved)
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

		prop[index] = convertValueType(val)
		updateYaml(yml, props, prop)
		return nil
	}

	if len(propsArr) == 1 {
		prop, _ := yml.Map()
		prop[path] = convertValueType(val)
		return nil
	}

	tmp, _ := get(yml, newPath)
	prop, err := tmp.Map()
	if err != nil {
		return err
	}

	prop[propName] = convertValueType(val)

	return nil
}

func SetValueForType(yaml *simpleyaml.Yaml, path string, value *simpleyaml.Yaml) error {
	if v, err := value.String(); err == nil {
		err := Set(yaml, path, v)
		return err
	}
	if v, err := value.Bool(); err == nil {
		err := Set(yaml, path, v)
		return err
	}
	if v, err := value.Int(); err == nil {
		err := Set(yaml, path, v)
		return err
	}
	if v, err := value.Float(); err == nil {
		err := Set(yaml, path, v)
		return err
	}
	if v, err := value.Array(); err == nil {
		err := Set(yaml, path, v)
		return err
	}
	if v, err := value.Map(); err == nil {
		err := Set(yaml, path, v)
		return err
	}
	return nil
}

func updateYaml(yml *simpleyaml.Yaml, props []string, prop []interface{}) {
	var yaml map[interface{}]interface{}
	propName := props[len(props)-1]

	if len(props) > 1 {
		props = props[:len(props)-1]
		tmp, _ := get(yml, strings.Join(props, "."))
		yaml, _ = tmp.Map()
	} else {
		yaml, _ = yml.Map()
	}

	yaml[propName] = prop
}
