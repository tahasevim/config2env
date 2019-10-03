package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//JsonParser is the type for parsing JSON files
type JsonParser struct{}

// Parse parses given JSON file and returns environment variables as key-value pairs
func (j *JsonParser) Parse(filename, prefix string) []EnvPair {
	fileContent := readFile(filename)
	result := j.getResultMap(fileContent)
	envs := []EnvPair{}
	for i, r := range result {
		if len(result) != 1 {
			j.collectEnvs(&envs, prefix+fmt.Sprintf("%d", i+1), r)
		} else {
			j.collectEnvs(&envs, prefix, r)
		}
	}
	return envs
}

// collectEnvs traverse JSON files recursively and generates environment variables
func (j *JsonParser) collectEnvs(envs *[]EnvPair, prefix string, mp map[string]interface{}) {
	for key, val := range mp {
		var newprefix string
		if prefix == "" {
			newprefix = strings.ToUpper(key)
		} else {
			newprefix = prefix + "_" + strings.ToUpper(key)
		}
		switch v := val.(type) {
		case map[string]interface{}:
			j.collectEnvs(envs, newprefix, v)
		case []interface{}:
			for i, elem := range v {
				arrPrefix := newprefix + "_" + strconv.Itoa(i+1)
				switch e := elem.(type) {
				case map[string]interface{}:
					j.collectEnvs(envs, arrPrefix, e)
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

// getResultMap simply loads read JSON file to a map and returns it
func (JsonParser) getResultMap(data []byte) []map[string]interface{} {
	var result []map[string]interface{} //Outher type may be array so it should be handled
	err := json.Unmarshal(data, &result)
	if err != nil {
		var temp map[string]interface{}
		err = json.Unmarshal(data, &temp)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, temp)
	}
	return result
}
