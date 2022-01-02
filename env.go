package env

import (
	"bufio"
	"os"
	"strings"
)

func Read(filepath string) (map[string]string, error) {
	env := make(map[string]string)
	// read the file line by line
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		env[line[0]] = line[1]
	}
	return env, nil
}
