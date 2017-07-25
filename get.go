package goml

import (
	"errors"
	"strconv"
	"strings"

	"github.com/smallfish/simpleyaml"
)

func GetFromFile(file string, path string) (interface{}, error) {
	yaml, err := ReadYamlFromFile(file)
	if err != nil {
		return nil, err
	}

	return Get(yaml, path)
}

func GetFromFileAsSimpleYaml(file string, path string) (*simpleyaml.Yaml, error) {
	yaml, err := ReadYamlFromFile(file)
	if err != nil {
		return nil, err
	}

	return GetAsSimpleYaml(yaml, path)
}

func Get(yml *simpleyaml.Yaml, path string) (interface{}, error) {
	val, ok := get(yml, path)
	if ok == nil {
		return nil, errors.New("property not found")
	}

	result, err := ExtractType(val)
	return result, err
}

func GetAsSimpleYaml(yml *simpleyaml.Yaml, path string) (*simpleyaml.Yaml, error) {
	val, ok := get(yml, path)
	if ok == nil {
		return nil, errors.New("property not found")
	}

	return val, nil
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
