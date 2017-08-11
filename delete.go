package goml

import (
	"errors"
	"os"
	"strconv"
	"strings"

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

func stringInSlice(element string, array []interface{}) (bool, int){
	for i := 0; i < len(array); i++ {
	 if array[i] == element {
		 return true, i
	 }

  }
	return false, 0
}

func Delete(yml *simpleyaml.Yaml, path string) error {
	propsArr := strings.Split(path, ".")
	propName := propsArr[len(propsArr)-1]
	props := propsArr[:len(propsArr)-1]
	newPath := strings.Join(props, ".")

	yaml_pointer, _ := get(yml, newPath)
	path_array, _ := yaml_pointer.Array()
	doesNameExist, my_index := stringInSlice(propName, path_array)

  if doesNameExist == true {
		tmp, props := get(yml, newPath)
		prop, err := tmp.Array()
		if err != nil {
			return err
		}
		prop = append(prop[:my_index], prop[my_index+1:]...)

		updateYaml(yml, props, prop)
		return nil
  }

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
