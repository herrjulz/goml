package goml

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/smallfish/simpleyaml"
	"gopkg.in/yaml.v2"
)

func Get(yml *simpleyaml.Yaml, path string) (interface{}, error) {
	val, ok := get(yml, path)
	if ok == nil {
		return nil, errors.New("property not found")
	}

	result, err := ExtractType(val)
	return result, err
}

func ExtractType(value *simpleyaml.Yaml) (interface{}, error) {
	if v, err := value.String(); err == nil {
		return v, nil
	}
	if v, err := value.Bool(); err == nil {
		return strconv.FormatBool(v), nil
	}
	if v, err := value.Int(); err == nil {
		return strconv.Itoa(v), nil
	}
	if v, err := value.Array(); err == nil {
		strSl := []string{}
		for _, val := range v {
			tmp := extractArrayType(val)
			strSl = append(strSl, tmp)
		}
		str := strings.Join(strSl, ",")
		return str, nil
	}
	if v, err := value.Map(); err == nil {
		return v, nil
	}
	return nil, errors.New("property not found")
}

func extractArrayType(value interface{}) string {
	switch t := value.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case string:
		return value.(string)
	case bool:
		return strconv.FormatBool(value.(bool))
	case int:
		return strconv.Itoa(value.(int))
	}
	return ""
}

func Set(yml *simpleyaml.Yaml, path string, val interface{}) error {
	propsArr := strings.Split(path, ".")
	propName := propsArr[len(propsArr)-1]
	props := propsArr[:len(propsArr)-1]
	newPath := strings.Join(props, ".")

	if index, err := strconv.Atoi(propName); err == nil {
		tmp, props := get(yml, newPath)
		if props == nil {
			return errors.New("peroperty not found")
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

func convertValueType(val interface{}) interface{} {
	switch val.(type) {
	default:
		return val
	case string:
		str := val.(string)
		if value, err := strconv.Atoi(str); err == nil {
			return value
		}
		if value, err := strconv.ParseBool(str); err == nil {
			return value
		}
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

	//val := yaml.MapSlice{}
	//err = yaml.Unmarshal([]byte(file), &val)
	//if err != nil {
	//return nil, errors.New("unmarshal []byte to yaml failed: " + err.Error())
	//}
	//fmt.Printf("--- m:\n%v\n\n", val)

	yml, _ := simpleyaml.NewYaml(file)
	return yml, nil
}

func get(yml *simpleyaml.Yaml, path string) (*simpleyaml.Yaml, []string) {
	solvedPath := []string{}

	props := strings.Split(path, ".")
	for _, p := range props {
		if index, err := strconv.Atoi(p); err == nil {
			yml = yml.GetIndex(index)
			solvedPath = append(solvedPath, strconv.Itoa(index))
			continue
		}

		if strings.Contains(p, ":") {
			if prop, err := yml.Array(); err == nil {
				index, err := returnIndexForProp(p, prop)
				if err != nil {
					return yml, nil
				}

				yml = yml.GetIndex(index)
				solvedPath = append(solvedPath, strconv.Itoa(index))
				continue
			}
		}
		solvedPath = append(solvedPath, p)
		yml = yml.Get(p)
	}
	return yml, solvedPath
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

func returnIndexForProp(propName string, array []interface{}) (int, error) {
	keyVal := strings.Split(propName, ":")
	key, val := keyVal[0], keyVal[1]

	for i, _ := range array {
		if key == "" {

			check := array[i]
			if check == val {
				return i, nil
			}

		} else {
			check := array[i].(map[interface{}]interface{})
			if check[key] == val {
				return i, nil
			}
		}
	}

	return 0, errors.New("property not found")
}
