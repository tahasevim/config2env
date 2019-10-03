package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// setupActual creates and returns a set manner map for generated environment variables
// This function is just for helper for testing stages
func setupActual(envs []EnvPair) map[string]struct{} {
	result := make(map[string]struct{})
	for _, env := range envs {
		result[fmt.Sprintf(env.Key+"="+env.Value)] = struct{}{}
	}
	return result
}

// setupExpected reads given env file and returns a set manner map for generated environment variables
// This function is just for helper for testing stages
func setupExpected(filename string) map[string]struct{} {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	fByte, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]struct{})
	for _, l := range strings.Split(string(fByte), "\n") {
		result[l] = struct{}{}
	}
	return result
}
