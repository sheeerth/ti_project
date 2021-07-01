package main

import (
	"errors"
	"fmt"
	_ "net"
	"os"
	_ "strconv"
	"strings"
)

func repeatedAddElements(stringArray []*string, actualString []string, index *int, master *IPAddress, level int) ([]string, *IPAddress,  error) {
	slave, _ := createNewIpAddress(strings.Replace(actualString[0], "\t", "", level), actualString[1])

	if master != nil {
		master.subnets = append(master.subnets, slave)
	}

	println(level, ":", actualString[0], index)

	*index++

	if *index >= len(stringArray) {
		return nil, nil, errors.New("out of range")
	}

	return strings.Split(*stringArray[*index], ","), slave, nil
}

func openFile(stringArray []*string) {
	var ipAddresses []*IPAddress
	var addElementError error

	index := 0

	actualString := strings.Split(*stringArray[index], ",")

	var master *IPAddress
	actualString, master, addElementError = repeatedAddElements(stringArray, actualString, &index, nil, 0)
	if addElementError != nil {
		return
	}

	if strings.Count(actualString[0], "\t") == 1 {
		for index <= len(stringArray) - 1 {
			var slave *IPAddress
			actualString, slave, addElementError = repeatedAddElements(stringArray, actualString, &index, master, 1)
			if addElementError != nil {
				break
			}

			if strings.Count(actualString[0], "\t") == 2 {
				for index <= len(stringArray) - 1 {

					actualString, _, addElementError = repeatedAddElements(stringArray, actualString, &index, slave, 2)
					if addElementError != nil {
						break
					}
				}
			}
		}
	}

	ipAddresses = append(ipAddresses, master)

	//var correctAddressList []IPAddress
	//
	//for x, address := range ipAddresses {
	//	repeat := false
	//
	//	for y, ipAddress := range ipAddresses {
	//		if x == y {
	//			break
	//		}
	//
	//		if address.ipAddress.IP.String() == ipAddress.ipAddress.IP.String() && address.ipAddress.Mask.String() == ipAddress.ipAddress.Mask.String() {
	//			repeat = true
	//		}
	//	}
	//
	//	if !repeat {
	//		correctAddressList = append(correctAddressList, ipAddresses[x])
	//	}
	//}

	fmt.Println("----")

	displayIpAddressArray(ipAddresses)
}

func fileOption(argsWithoutProg []string, index int) {
	stringArray, err := readFile(argsWithoutProg[index + 1])
	if err != nil {
		fmt.Println(err)
	}

	openFile(stringArray)
}

func main() {
	argsWithoutProg := os.Args[1:]

	for i, s := range argsWithoutProg {
		switch s {
		case "--file":
			fileOption(argsWithoutProg, i)
			break
		case "-f":
			fileOption(argsWithoutProg, i)
			break
		}
	}
}
