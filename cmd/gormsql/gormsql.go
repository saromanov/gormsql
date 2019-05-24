package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saromanov/gormsql/pkg/core"
	"github.com/saromanov/gormsql/pkg/sqlparser"
	"github.com/urfave/cli"
)

func createModelFromTables(path string) error {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read from file: %v", err)
	}

	tables, err := sqlparser.Parse(string(dat))
	if err != nil {
		return fmt.Errorf("unable to parse file: %v", err)
	}

	dir, err := getCurrentDirName()
	if err != nil {
		return fmt.Errorf("unable to get directory name: %v", err)
	}
	c := core.New(dir, "test", tables)
	if err := c.Do(); err != nil {
		return fmt.Errorf("unable to apply generation: %v", err)
	}
	return nil
}

// getCurrentDirName retruns current directory name
func getCurrentDirName() (string, error) {
	return os.Getwd()
}

func run(path string) error {
	return createModelFromTables(path)
}
func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "dir",
			Value: ".",
			Usage: "link to the dir for generated model",
		},
	}
	app.Name = "gormsql"
	app.Usage = "checking of availability of sites"
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "path to the dir or file",
			Action: func(c *cli.Context) error {
				modelPath := c.Args().First()
				dir := c.String("dir")
				fmt.Println(dir)
				if err := run(modelPath); err != nil {
					panic(err)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
