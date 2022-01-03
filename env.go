package env

import (
	"bufio"
	"os"
	"strings"
)

func Load() error {
	return LoadFrom(".env")
}

func LoadFrom(filepath string) error {
	// read the file line by line
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		os.Setenv(line[0], line[1])
	}
	return nil
}

func SetFrom(m map[string]string) {
	for k, v := range m {
		os.Setenv(k, v)
	}
}

func EnvAsMap(keys []string) map[string]string {
	env := make(map[string]string)
	for _, v := range keys {
		env[v] = os.Getenv(v)
	}
	return env
}

func GetEnvValsFromVarSlice(keys []string) []string {
	vals := []string{}
	for _, v := range keys {
		vals = append(vals, os.Getenv(v))
	}
	return vals
}
