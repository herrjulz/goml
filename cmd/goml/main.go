package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/JulzDiverse/goml"
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
	cmd.Version = "0.7.0"
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
				cli.BoolFlag{Name: "dry-run, d", Usage: "print set result to stdout"},
				cli.BoolFlag{Name: "json, j", Usage: "format output as json"},
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
		{
			Name:   "paths",
			Usage:  "Get paths of a yaml file",
			Action: getPaths,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file, f", Usage: "path to YAML file"},
			},
		},
	}

	err := cmd.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}

func getParam(c *cli.Context) {
	if c.NumFlags() != 4 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
		os.Exit(1)
	}

	res, err := goml.GetFromFile(c.String("file"), c.String("prop"))
	exitWithError(err)

	fmt.Println(res)
}

func setParam(c *cli.Context) {
	var key, value string
	keyVal := strings.Split(c.String("prop"), "=")
	key = keyVal[0]
	if len(keyVal) == 2 {
		value = keyVal[1]
	}
	if c.String("value") != "" {
		value = c.String("value")
	}

	var err error
	if c.String("key") != "" {
		err = goml.SetKeyInFile(c.String("file"), key, c.String("key"))
	} else {
		if c.Bool("dry-run") {
			bytes, err := ioutil.ReadFile(c.String("file"))
			exitWithError(err)
			output, err := goml.SetInMemory(bytes, key, value, c.Bool("json"))
			exitWithError(err)
			fmt.Println(string(output))
		} else {
			err = goml.SetInFile(c.String("file"), key, value)
		}
	}

	exitWithError(err)
}

func deleteParam(c *cli.Context) {
	if c.NumFlags() != 4 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

	err := goml.DeleteInFile(c.String("file"), c.String("prop"))
	exitWithError(err)
}

func transferParam(c *cli.Context) {
	if c.NumFlags() != 6 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

	err := goml.TransferToFile(c.String("file"), c.String("prop"), c.String("df"), c.String("dp"))
	exitWithError(err)
}

func getPaths(c *cli.Context) {
	filepath := c.String("file")
	file, err := ioutil.ReadFile(filepath)
	exitWithError(err)

	paths, err := goml.GetPaths(file)
	exitWithError(err)

	for _, path := range paths {
		fmt.Println(path)
	}
}

func exitWithError(err error) {
	if err != nil {
		r := color.New(color.FgHiRed)
		r.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
