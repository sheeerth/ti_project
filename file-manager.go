package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"ti/main/ipAddress"
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

func writeInFile(f *os.File, line string) error {
	_, err := f.WriteString(fmt.Sprintf("%s\n", line))

	if err != nil {
		return errors.New("something is wrong with file")
	}

	return nil
}

func saveAddressesInFile(addressesArray []*ipAddress.IPAddress) error {
	outputFile := createOutputFile("output.txt")

	err := ipAddress.SaveAddressInFile(addressesArray, 0, outputFile ,writeInFile)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	return nil
}
