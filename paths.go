package goml

import (
	"fmt"
	"strings"

	"github.com/smallfish/simpleyaml"
	"github.com/smallfish/simpleyaml/helper/util"
)

func GetPaths(file []byte) ([]string, error) {
	fmt.Println("INPUT:", string(file))
	yml, err := simpleyaml.NewYaml(file)
	if err != nil {
		return nil, err
	}
	paths, err := util.GetAllPaths(yml)
	if err != nil {
		return nil, err
	}
	fmt.Println("PATHS BEFORE:", paths)

	for i, _ := range paths {
		paths[i] = strings.Replace(paths[i], "/", ".", -1)
	}

	fmt.Println("RETURNING:", paths)
	return paths, nil
}
