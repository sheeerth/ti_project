package main

import (
	"bufio"
	"os"
)

func readFile(filePath string) ([]*string, error) {
	var stringArray []*string
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		stringArray = append(stringArray, &line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stringArray, nil
}
