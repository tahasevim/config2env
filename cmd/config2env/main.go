package main

import (
	"flag"

	"github.com/tahasevim/config2env/operator"
)

func main() {
	filename := flag.String("file", "test.yaml", "Name of file to be parsed")
	filetype := flag.String("type", "yaml", "Type of the given file")
	outputFile := flag.String("output", "config.env", "Path of output file")
	prefix := flag.String("prefix", "", "Prefix for output variables")
	flag.Parse()

	o := operator.NewOperator(*filename, *filetype, *outputFile, *prefix)
	o.Start()
	o.LogScreen()
	o.LogFile()
}
