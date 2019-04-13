package main

import (
	"flag"
	"fmt"

	"github.com/saromanov/gormsql/pkg/sqlparser"
	"github.com/saromanov/gormsql/pkg/os"
)

var (
	sqlTablesPath = flag.String("word", "foo", "a string")
)

func createModelFromTables(path string){
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = sqlparser.Parse(dat)
	if err != nil {
		panic(err)
	}
}
func run(){
	if *sqlTablesPath != "" {
		createModelFromTables(*sqlTablesPath)
	}
}
func main(){
	flag.Parse()
	run()
}