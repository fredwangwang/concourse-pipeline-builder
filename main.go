package main

import (
	"github.com/fredwangwang/concourse-pipeline-builder/commands"
	"github.com/jessevdk/go-flags"
	"log"
)

type Options struct {
	Import commands.Import `command:"import" description:"import the existing pipeline"`
}

func main() {
	var opt = Options{
		Import: commands.Import{},
	}

	parser := flags.NewParser(&opt, flags.HelpFlag)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
}
