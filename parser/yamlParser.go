package parser

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

//YamlParser is the type for parsing YAML files
type YamlParser struct{}

// Parse parses given YAML file and returns environment variables as key-value pairs
func (y *YamlParser) Parse(filename, prefix string) []EnvPair {
	fileContent := readFile(filename)
	result := y.getResultMap(fileContent)
	envs := make([]EnvPair, 0, 0)
	for i, r := range result {
		if len(result) != 1 {
			y.collectEnvs(&envs, prefix+fmt.Sprintf("%d", i+1), r)
		} else {
			y.collectEnvs(&envs, prefix, r)
		}
	}
	return envs
}

// collectEnvs traverse YAML files recursively and generates environment variables
func (y *YamlParser) collectEnvs(envs *[]EnvPair, prefix string, mp map[interface{}]interface{}) {
	for key, val := range mp {
		var newprefix string
		if prefix == "" {
			newprefix = strings.ToUpper(key.(string))
		} else {
			newprefix = prefix + "_" + strings.ToUpper(key.(string))
		}
		switch v := val.(type) {
		case map[interface{}]interface{}:
			y.collectEnvs(envs, newprefix, v)
		case []interface{}:
			for i, elem := range v {
				arrPrefix := newprefix + "_" + strconv.Itoa(i+1)
				switch e := elem.(type) {
				case map[interface{}]interface{}:
					y.collectEnvs(envs, arrPrefix, e)
				default:
					env := generateEnvPair(arrPrefix, e)
					*envs = append(*envs, *env)
				}
			}

		default:
			env := generateEnvPair(newprefix, v)
			*envs = append(*envs, *env)
		}
	}
}

// getResultMap simply loads read YAML file to a map and returns it
func (YamlParser) getResultMap(data []byte) []map[interface{}]interface{} {
	var result []map[interface{}]interface{}
	err := yaml.Unmarshal(data, &result)
	if err != nil {
		var temp map[interface{}]interface{}
		err = yaml.Unmarshal(data, &temp)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, temp)
	}
	return result
}
