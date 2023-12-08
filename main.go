package main

import (
	adventofcode "advent/src"
	"errors"
	"os"
)

func args(argv []string) (string, string, error) {
	if len(argv) < 3 {
		return "", "", errors.New("not enough arguments")
	}
	return argv[1], argv[2], nil
}

func main() {
	filename, ex, err := args(os.Args)

	if err != nil {
		return
	}
	adventofcode.Main(filename, ex)
}
