package main

import (
	"fmt"
	_ "net"
	"os"
	_ "strconv"
	"ti/main/ipAddress"
)

func fileOption(argsWithoutProg []string, index int) {
	stringArray, err := readFile(argsWithoutProg[index + 1])
	if err != nil {
		fmt.Println(err)
	}

	outputString := ipAddress.AddressesMapper(stringArray)
	//
	//cmd := exec.Command("stty", "size")
	//cmd.Stdin = os.Stdin
	//out, _ := cmd.Output()
	//fmt.Printf("out: %#v", string(out))

	println("")

	ipAddress.DisplayIpAddressArray(outputString, 0)
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
