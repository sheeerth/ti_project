package ipAddress

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

type IPAddress struct {
	IpAddress   net.IPNet
	description string
	subnets     []*IPAddress
}

func (x *IPAddress) giveInfoString() (string, error) {
	if x == nil {
		return "", errors.New("object can't be nil")
	}

	return fmt.Sprintf("%s %s %s", (*x).IpAddress.IP.To4(), (*x).IpAddress.Mask.String(), (*x).description), nil
}

func (x *IPAddress) getAddressString () string {
	return x.IpAddress.IP.String()
}

func (x *IPAddress) getAddressMaskString() string {
	return x.IpAddress.Mask.String()
}

func CreateNewIpAddress(addressString string, description string) (*IPAddress, error) {
	_, ipv4Net, err := net.ParseCIDR(addressString)
	if err != nil {
		return nil, errors.New("wrong")
	}
	var test []*IPAddress = nil

	return &IPAddress{IpAddress: *ipv4Net, description: description, subnets: test}, nil
}

func SaveAddressInFile(addressArray []*IPAddress, level int, stream *os.File,display func(stream *os.File, a string) error) error {
	for _, address := range addressArray {
		outputString, err := address.giveInfoString()
		if err != nil {
			fmt.Println(err)
		}

		var tabulationArray []string

		for i := 0; i < level; i++ {
			tabulationArray = append(tabulationArray, "\t")
		}

		saveIsWrong := display(stream, strings.Join([]string{strings.Join(tabulationArray, ""), outputString}, ""))

		if saveIsWrong != nil {
			return saveIsWrong
		}

		if address.subnets != nil {
			err2 := SaveAddressInFile(address.subnets, level + 1, stream ,display)

			if err2 != nil {
				return err2
			}
		}
	}

	return nil
}