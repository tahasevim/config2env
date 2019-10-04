package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/tahasevim/config2env/operator"
)

func main() {
	input := flag.String("input", "", "Name of file to be parsed")
	output := flag.String("output", "", "Path of output file")
	filetype := flag.String("type", "", "Type of the given file")
	prefix := flag.String("prefix", "", "Prefix for output variables") //Optional
	flag.Parse()

	if *input == "" {
		fmt.Println((errors.New("Input file is missing.\nusage: config2env --input config.json --type json --output config.env")))
		return
	}

	if *output == "" {
		fmt.Println((errors.New("Output file is missing.\nusage: config2env --input config.json --type json --output config.env")))
		return
	}

	if *filetype == "" {
		fmt.Println((errors.New("File type is missing.\nusage: config2env --input config.json --type json --output config.env")))
		return
	}

	o := operator.NewOperator(*input, *filetype, *output, *prefix)
	o.Start()
	o.LogScreen()
	o.LogFile()
}
