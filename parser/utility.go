package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// generateEnvPair creates environment variables when parser reaches the maximum depth
// Maximum depth means that the value is a primitive type so it is not nested anymore
func generateEnvPair(prefix string, value interface{}) *EnvPair {
	return NewEnvPair(prefix, fmt.Sprintf("%v", value))
}

// readFile reads given file
func readFile(filename string) []byte {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	fByte, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return fByte
}
