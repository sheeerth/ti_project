package main

import (
	"errors"
	"fmt"
	"log"
	_ "net"
	"os"
	_ "strconv"
	"ti/main/ipAddress"
)

func fileOption(argsWithoutProg []string, index int) []*ipAddress.IPAddress {
	stringArray, err := readFile(argsWithoutProg[index + 1])
	if err != nil {
		fmt.Println(err)
	}

	return ipAddress.AddressesMapper(stringArray)
}

func createOutputFile(name string) *os.File {
	f, err := os.Create(name)

	if err != nil {
		log.Fatal(err)
	}

	return f
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

func main() {
	argsWithoutProg := os.Args[1:]

	for i, s := range argsWithoutProg {
		switch s {
		case "--file":
			mainAddressesArray := fileOption(argsWithoutProg, i)

			err := saveAddressesInFile(mainAddressesArray)

			if err != nil {
				println(err.Error())
			}
			break
		case "-f":
			mainAddressesArray := fileOption(argsWithoutProg, i)

			err := saveAddressesInFile(mainAddressesArray)

			if err != nil {
				println(err.Error())
			}
			break
		}
	}
}
