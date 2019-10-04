package parser

import (
	"reflect"
	"testing"
)

func TestYamlParse(t *testing.T) {
	tests := []struct {
		yamlFile string
		envFile  string
	}{
		{
			yamlFile: "../test/yaml/test1.yaml",
			envFile:  "../test/yaml/test1.env",
		},
		{
			yamlFile: "../test/yaml/test2.yaml",
			envFile:  "../test/yaml/test2.env",
		},
	}

	yamlParser := YamlParser{}
	for _, test := range tests {
		actual := setupActual(yamlParser.Parse(test.yamlFile, ""))
		expected := setupExpected(test.envFile)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\nActual:\n%v\nExpected:\n%v", actual, expected)
		}
	}
}
