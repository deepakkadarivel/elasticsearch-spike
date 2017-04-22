package config

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {

		panic(err)
		os.Exit(1)
	}
	return string(file)
}
