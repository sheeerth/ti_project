package ipAddress

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type IPAddress struct {
	ipAddress   net.IPNet
	description string
	subnets     []*IPAddress
}

func giveInfoString(x *IPAddress) (string, error) {
	if x == nil {
		return "", errors.New("object can't be nil")
	}

	return fmt.Sprintf("%s %s %s", (*x).ipAddress.IP.To4(), (*x).ipAddress.Mask.String(), (*x).description), nil
}

func CreateNewIpAddress(addressString string, description string) (*IPAddress, error) {
	_, ipv4Net, err := net.ParseCIDR(addressString)
	if err != nil {
		return nil, errors.New("wrong")
	}
	var test []*IPAddress = nil

	return &IPAddress{ipAddress: *ipv4Net, description: description, subnets: test}, nil
}

func DisplayIpAddressArray(addressArray []*IPAddress, level int) {
	for _, address := range addressArray {
		outputString, err := giveInfoString(address)
		if err != nil {
			fmt.Println(err)
		}

		var tabulationArray []string

		for i := 0; i < level; i++ {
			tabulationArray = append(tabulationArray, "\t")
		}

		fmt.Println(strings.Join(tabulationArray, ""), outputString)

		if address.subnets != nil {
			DisplayIpAddressArray(address.subnets, level + 1)
		}
	}
}