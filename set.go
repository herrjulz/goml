package goml

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/smallfish/simpleyaml"
)

const (
	ADDMAP   = "GOML_ADD_MAP"
	ADDARRAY = "GOML_ADD_ARRAY"
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

func SetInMemory(file []byte, path string, val interface{}, asJson bool) ([]byte, error) {
	if len(file) == 0 {
		file = append(file, []byte(`{}`)...)
	}

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

	var ymlBytes []byte
	if asJson {
		ymlBytes, err = json.Marshal(ymlMap)
		if err != nil {
			return nil, err
		}
	} else {
		ymlBytes, err = yaml.Marshal(ymlMap)
		if err != nil {
			return nil, err
		}
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
			if err = Set(yml, newPath, ADDARRAY); err != nil {
				return err //errors.New("SORRY CANNOT DO THAT. property not found")
			}
			tmp, props = get(yml, newPath)
		}

		prop, err := tmp.Array()
		if err != nil {
			return err
		}

		prop[index] = convertValueType(setValueOrAddChild(val))

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

	if strings.Contains(propName, ":") || strings.Contains(propName, "|") {
		tmp, props := get(yml, newPath)
		prop, err := tmp.Array()
		if err != nil {
			if err = Set(yml, newPath, ADDARRAY); err != nil {
				return err //errors.New("SORRY CANNOT DO THAT. property not found")
			}
			tmp, props = get(yml, newPath)
			prop, err = tmp.Array()
		}

		index, err := returnIndexForProp(propName, prop)
		if err != nil {
			index = createArrayEntry(propName, &prop)
		} else {
			prop[index] = convertValueType(setValueOrAddChild(val))
		}

		updateYaml(yml, props, prop)
		return nil
	}

	if len(propsArr) == 1 {
		prop, _ := yml.Map()
		prop[path] = convertValueType(setValueOrAddChild(val))
		return nil
	}

	tmp, _ := get(yml, newPath)
	prop, err := tmp.Map()
	if err != nil {
		if err = Set(yml, newPath, ADDMAP); err != nil {
			return err
		}
		tmp, _ = get(yml, newPath)
		prop, err = tmp.Map()
	}

	prop[propName] = convertValueType(setValueOrAddChild(val))

	return nil
}

func setValueOrAddChild(val interface{}) interface{} {
	switch val {
	case ADDMAP:
		m := make(map[interface{}]interface{})
		return m
	case ADDARRAY:
		return []interface{}{}
	}
	return val
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
