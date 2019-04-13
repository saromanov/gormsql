package main

import (
	"flag"

	"github.com/saromanov/gormsql/pkg/os"
	"github.com/saromanov/gormsql/pkg/sqlparser"
)

var (
	sqlTablesPath = flag.String("sql", "", "Path to the file .sql")
)

func createModelFromTables(path string) {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = sqlparser.Parse(dat)
	if err != nil {
		panic(err)
	}
}
func run() {
	if *sqlTablesPath != "" {
		createModelFromTables(*sqlTablesPath)
	}
}
func main() {
	flag.Parse()
	run()
}
