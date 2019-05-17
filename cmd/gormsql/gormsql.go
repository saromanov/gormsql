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

	c := core.New("test", tables)
	if err := c.Do(); err != nil {
		return fmt.Errorf("unable to apply generation: %v", err)
	}
	return nil
}
func run(path string) error {
	return createModelFromTables(path)
}
func main() {
	app := cli.NewApp()
	app.Name = "gormsql"
	app.Usage = "checking of availability of sites"
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "path to the dir or file",
			Action: func(c *cli.Context) error {
				modelPath := c.Args().First()
				fmt.Println("MODEL: ", modelPath)
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
