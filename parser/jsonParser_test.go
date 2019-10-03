package parser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		jsonFile string
		envFile  string
	}{
		{
			jsonFile: "../test/json/test1.json",
			envFile:  "../test/json/test1.env",
		},
		{
			jsonFile: "../test/json/test2.json",
			envFile:  "../test/json/test2.env",
		},
	}

	jsParser := JsonParser{}
	for _, test := range tests {
		actual := setupActual(jsParser.Parse(test.jsonFile, ""))
		expected := setupExpected(test.envFile)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\nActual:\n%v\nExpected:\n%v", actual, expected)
		}
	}
}
