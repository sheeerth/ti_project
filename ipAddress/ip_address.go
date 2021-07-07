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
	Description string
	Subnets     []*IPAddress
}

func (x *IPAddress) GiveInfoString() (string, error) {
	if x == nil {
		return "", errors.New("object can't be nil")
	}

	onesMask, _ := x.IpAddress.Mask.Size()

	return fmt.Sprintf("%s/%d %s", (*x).IpAddress.IP.To4(), onesMask, (*x).Description), nil
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

	return &IPAddress{IpAddress: *ipv4Net, Description: description, Subnets: test}, nil
}

func SaveAddressInFile(addressArray []*IPAddress, level int, stream *os.File,display func(stream *os.File, a string) error) error {
	for _, address := range addressArray {
		outputString, err := address.GiveInfoString()
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

		if address.Subnets != nil {
			err2 := SaveAddressInFile(address.Subnets, level + 1, stream ,display)

			if err2 != nil {
				return err2
			}
		}
	}

	return nil
}

func (x *IPAddress) SubnetsContainsNetwork(address string) bool {
	for _, subnet := range x.Subnets {
		if subnet.IpAddress.IP.String() == address {
			return true
		}
	}

	return false
}