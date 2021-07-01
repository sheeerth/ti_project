package main

import (
	"errors"
	"fmt"
	"net"
)

type IPAddress struct {
	ipAddress net.IPNet
	description string
	subnets []*IPAddress
}

func giveInfoString(x *IPAddress) (string, error) {
	if x == nil {
		return "", errors.New("object can't be nil")
	}

	return fmt.Sprintf("%s %s %s", (*x).ipAddress.IP.To4(), (*x).ipAddress.Mask.String(), (*x).description), nil
}

func createNewIpAddress(addressString string, description string) (*IPAddress, error) {
	_, ipv4Net, err := net.ParseCIDR(addressString)
	if err != nil {
		return nil, errors.New("Wrong")
	}
	var test []*IPAddress = nil

	return &IPAddress{ipAddress: *ipv4Net, description: description, subnets: test}, nil
}

func displayIpAddressArray(addressArray []*IPAddress) {
	for _, address := range addressArray {
		outputString, err := giveInfoString(address)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(outputString)

		if address.subnets != nil {
			for _, subnet := range address.subnets {
				outputString, err := giveInfoString(subnet)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("\t", outputString)

				if subnet.subnets != nil {
					for _, nextLvl := range subnet.subnets {
						outputString, err := giveInfoString(nextLvl)
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println("\t\t", outputString)
					}
				}
			}
		}
	}
}