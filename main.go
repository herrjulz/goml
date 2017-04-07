package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JulzDiverse/goml/goml"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func main() {
	cmd := cli.NewApp()
	cmd.Authors = []cli.Author{
		cli.Author{
			Name:  "Julz Skupnjak",
			Email: "julian.skupnjak@gmail.com",
		},
	}
	cmd.Name = "goml"
	cmd.Usage = "CLI Tool to do CRUD like manipulation on YAML files"
	cmd.Version = "0.1.0"
	cmd.Commands = []cli.Command{
		{
			Name:   "get",
			Usage:  "Get property",
			Action: getParam,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file, f", Usage: "path to YAML file"},
				cli.StringFlag{Name: "prop, p", Usage: "property path string - foo.bar.zoo"},
			},
		},
		{
			Name:   "set",
			Usage:  "Set/Update property",
			Action: setParam,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file, f", Usage: "path to YAML file"},
				cli.StringFlag{Name: "prop, p", Usage: "property path string - foo.bar.zoo"},
				cli.StringFlag{Name: "value, v", Usage: "value for the defined property"},
				cli.StringFlag{Name: "key, k", Usage: "private key file"},
			},
		},
		{
			Name:   "delete",
			Usage:  "Delete property",
			Action: deleteParam,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file, f", Usage: "path to YAML file"},
				cli.StringFlag{Name: "prop, p", Usage: "property path string - foo.bar.zoo"},
			},
		},
		{
			Name:   "transfer",
			Usage:  "Transfer property",
			Action: transferParam,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file, f", Usage: "path to YAML file"},
				cli.StringFlag{Name: "prop, p", Usage: "property path (string) - foo.bar.zoo"},
				cli.StringFlag{Name: "df", Usage: "destination YAML file"},
				cli.StringFlag{Name: "dp", Usage: "destination property path (string) - foo.bar.zoo"},
			},
		},
	}
	cmd.Run(os.Args)
}

func getParam(c *cli.Context) {
	if c.NumFlags() != 4 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
		os.Exit(1)
	}

	yaml, err := goml.ReadYamlFromFile(c.String("file"))
	exitWithError(err)

	res, err := goml.Get(yaml, c.String("prop"))
	exitWithError(err)

	fmt.Println(res)
}

func setParam(c *cli.Context) {
	if c.NumFlags() != 6 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

	yaml, err := goml.ReadYamlFromFile(c.String("file"))
	exitWithError(err)

	var value string
	if c.String("value") != "" {
		value = c.String("value")
	} else if c.String("key") != "" {
		bytes, err := ioutil.ReadFile(c.String("key"))
		exitWithError(err)
		value = string(bytes)
	}

	err = goml.Set(yaml, c.String("prop"), value)
	exitWithError(err)
	goml.WriteYaml(yaml, c.String("file"))
}

func deleteParam(c *cli.Context) {
	if c.NumFlags() != 4 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

	yaml, err := goml.ReadYamlFromFile(c.String("file"))
	exitWithError(err)

	err = goml.Delete(yaml, c.String("prop"))
	if err != nil {
		exitWithError(errors.New("Couldn't delete property for path: " + c.String("prop")))
	}

	goml.WriteYaml(yaml, c.String("file"))
}

func transferParam(c *cli.Context) {
	if c.NumFlags() != 6 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

	sourceYaml, err := goml.ReadYamlFromFile(c.String("file"))
	exitWithError(err)

	destYaml, err := goml.ReadYamlFromFile(c.String("df"))
	exitWithError(err)

	value, _ := goml.Get(sourceYaml, c.String("prop"))

	err = goml.Set(destYaml, c.String("dp"), value)
	exitWithError(err)

	goml.WriteYaml(destYaml, c.String("df"))
}

func exitWithError(err error) {
	if err != nil {
		r := color.New(color.FgHiRed)
		r.Println(err)
		os.Exit(1)
	}
}
