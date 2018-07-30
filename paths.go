package goml

import (
	"strings"

	"github.com/smallfish/simpleyaml"
	"github.com/smallfish/simpleyaml/helper/util"
)

func GetPaths(file []byte) ([]string, error) {
	yml, err := simpleyaml.NewYaml(file)
	if err != nil {
		return nil, err
	}
	paths, err := util.GetAllPaths(yml)
	if err != nil {
		return nil, err
	}

	for i, _ := range paths {
		paths[i] = strings.Replace(paths[i], "/", ".", -1)
	}

	return paths, nil
}
