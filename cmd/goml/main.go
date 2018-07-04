package main

import (
	"errors"
	"fmt"
	"os"

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
	cmd.Version = "0.4.0"
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

	res, err := goml.GetFromFile(c.String("file"), c.String("prop"))
	exitWithError(err)

	fmt.Println(res)
}

func setParam(c *cli.Context) {
	if c.NumFlags() != 6 {
		cli.ShowAppHelp(c)
		exitWithError(errors.New("invalid number of arguments"))
	}

	var err error
	if c.String("value") != "" {
		value := c.String("value")
		err = goml.SetInFile(c.String("file"), c.String("prop"), value)
	} else if c.String("key") != "" {
		err = goml.SetKeyInFile(c.String("file"), c.String("prop"), c.String("key"))
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

func exitWithError(err error) {
	if err != nil {
		r := color.New(color.FgHiRed)
		r.Println(err)
		os.Exit(1)
	}
}
