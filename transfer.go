package goml

import "github.com/smallfish/simpleyaml"

func TransferToFile(from string, fromPath string, to string, toPath string) error {
	fromYaml, err := ReadYamlFromFile(from)
	if err != nil {
		return err
	}

	toYaml, err := ReadYamlFromFile(to)
	if err != nil {
		return err
	}

	err = Transfer(fromYaml, fromPath, toYaml, toPath)
	if err != nil {
		return err
	}

	return WriteYaml(toYaml, to)
}

func Transfer(fromYaml *simpleyaml.Yaml, fromPath string, toYaml *simpleyaml.Yaml, toPath string) error {
	value, err := Get(fromYaml, fromPath)
	if err != nil {
		return err
	}

	return Set(toYaml, toPath, value)
}
