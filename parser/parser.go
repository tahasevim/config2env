package parser

//Parser is main interface for that Operator type needs to parse given config files
//New config file types that will be parsed have to implement this method. (such as TOML and HCL configuration types)
type Parser interface {
	Parse(string, string) []EnvPair
}

//EnvPair is the representation of environment variables as key-value pairs
type EnvPair struct {
	Key   string
	Value string
}

//NewEnvPair returns a new instance of EnvPair
func NewEnvPair(key, value string) *EnvPair {
	return &EnvPair{
		Key:   key,
		Value: value,
	}
}
