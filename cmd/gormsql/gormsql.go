package main

import (
	"flag"
	"fmt"
)

var (
	sqlTablesPath = flag.String("word", "foo", "a string")
)

func run(){
	if *sqlTablesPath != "" {

	}
}
func main(){
	flag.Parse()
	run()
}