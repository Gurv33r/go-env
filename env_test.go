package env_test

import (
	"os"
	"strings"
	"testing"

	env "github.com/Gurv33r/go-env"
)

var fuzzMap map[string]string = make(map[string]string)

func TestLoadFrom(t *testing.T) {
	// create env file
	file, err := os.Create("test.env")
	if err != nil {
		panic(err)
	}
	// write some test data
	const fuzz = "A=1\nB=4\nC=HELLOWORLD"
	err = os.WriteFile("test.env", []byte(fuzz), os.ModeAppend)
	if err != nil {
		panic(err)
	}
	file.Close()
	if err = env.LoadFrom("test.env"); err != nil {
		t.Error(err)
	}
	if strings.Compare(os.Getenv("A"), "1") != 0 {
		t.Fail()
	}
	if strings.Compare(os.Getenv("B"), "4") != 0 {
		t.Fail()
	}
	if strings.Compare(os.Getenv("C"), "HELLOWORLD") != 0 {
		t.Fail()
	}
	os.Remove("test.env")
}

func TestSetFrom(t *testing.T) {
	fuzzMap["A"] = "1"
	fuzzMap["1"] = "?"
	fuzzMap["C"] = "Hello, world"
	env.SetFrom(fuzzMap)
	for k, v := range fuzzMap {
		if strings.Compare(os.Getenv(k), v) != 0 {
			t.Fail()
		}
	}
}
