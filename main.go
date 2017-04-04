package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/JulzDiverse/goml/goml"
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
	cmd.Version = "0.0.1"
	cmd.Commands = []cli.Command{
		{
			Name:      "get",
			Usage:     "Get property",
			ArgsUsage: "foo.bar.zoo",
			Action:    getParam,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file, f", Usage: "path to YAML file"},
				cli.StringFlag{Name: "prop, p", Usage: "property path string - foo.bar.zoo"},
			},
		},
		{
			Name:      "set",
			Usage:     "Set property",
			ArgsUsage: "get foo.bar.zoo ",
			Action:    setParam,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file, f", Usage: "path to YAML file"},
				cli.StringFlag{Name: "prop, p", Usage: "property path string - foo.bar.zoo"},
				cli.StringFlag{Name: "value, v", Usage: "set the value for the defined property"},
			},
		},
		{
			Name:      "delete",
			Usage:     "delete property",
			ArgsUsage: "delete foo.bar.zoo ",
			Action:    deleteParam,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file, f", Usage: "path to YAML file"},
				cli.StringFlag{Name: "prop, p", Usage: "property path string - foo.bar.zoo"},
			},
		},
		{
			Name:      "transfer",
			Usage:     "transfer property",
			ArgsUsage: "transfer",
			Action:    transferParam,
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

	rawValue, _ := goml.Get(yaml, c.String("prop"))
	// exitWithError(err)

	if rawValue == nil {
		exitWithError(errors.New("Couldn't find property"))
	}

	res, err := goml.ExtractType(rawValue)
	exitWithError(err)

	// fmt.Printf("%s", rawValue)
	fmt.Println(res)
}

func setParam(c *cli.Context) {
	if c.NumFlags() != 6 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

	yaml, err := goml.ReadYamlFromFile(c.String("file"))
	exitWithError(err)

	err = goml.Set(yaml, c.String("prop"), c.String("value"))
	exitWithError(err)

	goml.WriteYaml(yaml, c.String("file"))
}

func deleteParam(c *cli.Context) {
	if c.NumFlags() != 4 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
		os.Exit(1)
	}

	yaml, err := goml.ReadYamlFromFile(c.String("file"))
	exitWithError(err)

	updatedYaml, err := goml.Delete(yaml, c.String("prop"))
	exitWithError(err)

	goml.WriteYaml(updatedYaml, c.String("file"))
}

func transferParam(c *cli.Context) {
	if c.NumFlags() != 6 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
		os.Exit(1)
	}

	sourceYaml, err := goml.ReadYamlFromFile(c.String("file"))
	exitWithError(err)

	destYaml, err := goml.ReadYamlFromFile(c.String("df"))
	exitWithError(err)

	value, _ := goml.Get(sourceYaml, c.String("prop"))

	err = goml.SetValueForType(destYaml, c.String("dp"), value)
	exitWithError(err)

	goml.WriteYaml(destYaml, c.String("df"))
}

func exitWithError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
