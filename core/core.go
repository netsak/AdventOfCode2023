package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FailOnError prints a message and panics
// if there is an error.
func FailOnError(err error, msg string) {
	if err != nil {
		fmt.Println("ERROR:", msg)
		panic(err)
	}
}

// Assert that the two values are equal.
// If not both are printed for comparison.
func Assert[T comparable](expected T, actual T) bool {
	if expected == actual {
		return true
	}
	fmt.Printf("ASSERTATION ERROR !!!!!!!!!!\nexpected: %+v\nactual:   %+v\n", expected, actual)
	return false
}

// ReadLines reads all not empty lines from a file
// with stripped leading and trailing whitespaces.
func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", filename, err)
	}
	defer f.Close()
	ret := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		ret = append(ret, line)
	}
	return ret, nil
}
