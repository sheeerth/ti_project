package main

import (
	"fmt"
	"log"
	_ "net"
	"os"
	"sort"
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

func main() {
	argsWithoutProg := os.Args[1:]

	for i, s := range argsWithoutProg {
		switch s {
		case "--file":
			mainAddressesArray := fileOption(argsWithoutProg, i)

			for _, address := range mainAddressesArray {
				address.ExtendSubnet()

				sort.Sort(address.Subnets)

				address.Subnets = address.Subnets.ReduceMaskOfSubnets()
				sort.Sort(address.Subnets)

				address.Subnets = address.Subnets.ReduceMaskOfSubnets()
				sort.Sort(address.Subnets)

				address.Subnets = address.Subnets.ReduceMaskOfSubnets()
				sort.Sort(address.Subnets)

				address.Subnets = address.Subnets.ReduceMaskOfSubnets()
				sort.Sort(address.Subnets)

				address.Subnets = address.Subnets.ReduceMaskOfSubnets()
				sort.Sort(address.Subnets)

				address.Subnets = address.Subnets.ReduceMaskOfSubnets()
				sort.Sort(address.Subnets)

				// TODO do pętli i jakis mechanizm wykrywania kiedy koniec
			}

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
