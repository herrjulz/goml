package goml

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/smallfish/simpleyaml"
	"gopkg.in/yaml.v2"
)

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
	if v, err := value.Float(); err == nil {
		return fmt.Sprint(v), nil
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
		fmt.Fprintf(os.Stderr, "unexpected type %T\n", t)
	case string:
		return value.(string)
	case bool:
		return strconv.FormatBool(value.(bool))
	case int:
		return strconv.Itoa(value.(int))
	case float64:
		return fmt.Sprint(value.(float64))
	}
	return ""
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

func WriteYaml(yml *simpleyaml.Yaml, file string) error {
	var (
		err  error
		goml interface{}
	)
	switch {
	case yml.IsMap():
		goml, err = yml.Map()
	case yml.IsArray():
		goml, err = yml.Array()
	default:
		err = errors.New("Unexpected yml structure after manipulation")
	}
	if err != nil {
		return err
	}

	gomlSave, err := yaml.Marshal(goml)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, gomlSave, 0644)
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

	if len(file) == 0 {
		file = append(file, []byte(`{}`)...)
	}

	return simpleyaml.NewYaml(file)
}

func returnIndexForProp(propName string, array []interface{}) (int, error) {
	keyVal := strings.Split(propName, ":")
	if len(keyVal) < 2 {
		keyVal = strings.Split(propName, "|")
	}

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

func createArrayEntry(propName string, array *[]interface{}) int {
	keyVal := strings.Split(propName, ":")
	if len(keyVal) < 2 {
		keyVal = strings.Split(propName, "|")
	}

	key, val := keyVal[0], keyVal[1]
	m := make(map[interface{}]interface{})
	m[key] = val
	*array = append(*array, m)

	return len(*array) - 1
}
