package operator

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/tahasevim/config2env/parser"
)

// Operator is the main type of the tool that operates parsing and logging
type Operator struct {
	input    string
	filetype string
	output   string
	prefix   string
	parser.Parser
	envs []parser.EnvPair
}

// NewOperator creates and returns a new instance of the Operator type
func NewOperator(input, filetype, output, prefix string) *Operator {
	parsers := map[string]parser.Parser{
		"json": &parser.JsonParser{},
		"yaml": &parser.YamlParser{},
	}
	return &Operator{
		input:    input,
		filetype: filetype,
		output:   output,
		prefix:   strings.ToUpper(prefix),
		Parser:   parsers[filetype],
	}
}

// Start simply starts parsing operation
func (o *Operator) Start() {
	o.envs = o.Parse(o.input, o.prefix)
}

// LogFile writes generated environment variables to the output file.
func (o *Operator) LogFile() {
	f, err := os.Create(o.output)
	if err != nil {
		log.Fatalln(err)
	}

	for _, env := range o.envs {
		f.WriteString(env.Key + "=" + env.Value + "\n")
	}
}

// LogScreen prints generated environment variables to the screen as blue-yellow strings.
func (o *Operator) LogScreen() {
	for _, env := range o.envs {
		k := color.BlueString(env.Key)
		e := color.WhiteString("=")
		v := color.YellowString(env.Value)
		fmt.Println(k + e + v)
	}
}
