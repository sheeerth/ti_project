package ipAddress

import (
	"errors"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"
)

type ByIpAddress []*IPAddress

type IPAddress struct {
	ipAddress   net.IPNet
	description string
	subnets     ByIpAddress
}

func (x *IPAddress) toString() (string, error) {
	if x == nil { return "", errors.New("object can't be nil") }

	onesMask, _ := x.ipAddress.Mask.Size()

	return fmt.Sprintf("%s/%d %s", (*x).ipAddress.IP.To4(), onesMask, (*x).description), nil
}

func (x *IPAddress) getAddressString () string {
	return x.ipAddress.IP.String()
}

func (x *IPAddress) getAddressMaskString() string {
	return x.ipAddress.Mask.String()
}

func (x *IPAddress) SubnetsContainsNetwork(address string) bool {
	for _, subnet := range x.subnets {
		if subnet.ipAddress.IP.String() == address {
			return true
		}
	}

	return false
}

func (x *IPAddress) ExtendAndReduceSubnets() {
	x.ExtendSubnet()
	sort.Sort(x.subnets)

	subnetsLong := len(x.subnets)

	x.subnets = x.subnets.ReduceMaskOfSubnets()
	sort.Sort(x.subnets)

	afterSubnetLong := len(x.subnets)

	for subnetsLong != afterSubnetLong {
		subnetsLong = afterSubnetLong

		x.subnets = x.subnets.ReduceMaskOfSubnets()
		sort.Sort(x.subnets)

		afterSubnetLong = len(x.subnets)
	}

	for _, subnet := range x.subnets {
		if subnet.subnets != nil {
			subnet.ExtendAndReduceSubnets()
		}
	}
}

func (a ByIpAddress) Len() int           { return len(a) }
func (a ByIpAddress) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIpAddress) Less(i, j int) bool { return a[i].ipAddress.IP[2] < a[j].ipAddress.IP[2] }

func (a ByIpAddress) ReduceMaskOfSubnets() []*IPAddress  {
	calculate := 0
	var optimizeArray []*IPAddress

	for i := len(a) - 1; i >= 0; i-- {
		if calculate == 0 && i == 0 {
			optimizeArray = append(optimizeArray, a[i])
			break
		}

		if a[i].description == "WOLNA" {
			if i + 1 < len(a) - 1 && a[i - 1].ipAddress.Mask.String() != a[i].ipAddress.Mask.String() {
				continue
			}

			calculate++
		} else {
			if calculate == 1 {
				optimizeArray = append(optimizeArray, a[i + 1])
			}

			optimizeArray = append(optimizeArray, a[i])
			calculate = 0
		}

		if calculate / 2 == 1 && (a[i].ipAddress.Mask.String() == a[i + 1].ipAddress.Mask.String()) {
			for i2, b := range a[i].ipAddress.Mask {
				if b == 0x00 {
					a[i].ipAddress.Mask[i2 - 1] = a[i].ipAddress.Mask[i2 - 1]<<1
				}
			}

			optimizeArray = append(optimizeArray, a[i])
			calculate = 0
		} else if  calculate == 2 && i + 1 < len(a) - 1 && a[i].ipAddress.Mask.String() != a[i + 1].ipAddress.Mask.String(){
			optimizeArray = append(optimizeArray, a[i + 1])
			optimizeArray = append(optimizeArray, a[i])
			calculate = 0
		}
	}

	return optimizeArray
}

func (x *IPAddress) ExtendSubnet() {
	var subnetsAddresses [][]string
	var repeatList []string
	var correctAddressList []string

	onesMask, _ := x.ipAddress.Mask.Size()
	masterSubnets, _ := hosts(fmt.Sprintf("%s/%d", x.ipAddress.IP.String(), onesMask))

	for _, subnet := range x.subnets {
		onesMask, _ := subnet.ipAddress.Mask.Size()
		subnetSubnets, _ := hosts(fmt.Sprintf("%s/%d", subnet.ipAddress.IP.String(), onesMask))

		subnetsAddresses = append(subnetsAddresses, subnetSubnets)
	}

	for _, subnet := range masterSubnets {
		for _, subnetsAddress := range subnetsAddresses {
			for _, s2 := range subnetsAddress {
				if subnet == s2 {
					repeatList = append(repeatList, subnet)
				}
			}
		}
	}

	for _, subnet := range masterSubnets {
		if !arrayContainsElement(repeatList, subnet) {
			correctAddressList = append(correctAddressList, subnet)
		}
	}

	for _, s2 := range correctAddressList {
		addressString := fmt.Sprintf("%s/%d", s2, 24)
		slave, _ := CreateNewIpAddress(addressString, "WOLNA")

		x.subnets = append(x.subnets, slave)
	}
}

func CreateNewIpAddress(addressString string, description string) (*IPAddress, error) {
	_, ipv4Net, err := net.ParseCIDR(addressString)
	if err != nil { return nil, errors.New("wrong") }

	var test []*IPAddress = nil

	return &IPAddress{ipAddress: *ipv4Net, description: description, subnets: test}, nil
}

func SaveAddressInFile(addressArray []*IPAddress, level int, stream *os.File,display func(stream *os.File, a string) error) error {
	for _, address := range addressArray {
		outputString, err := address.toString()
		if err != nil { fmt.Println(err) }

		var tabulationArray []string

		for i := 0; i < level; i++ {
			tabulationArray = append(tabulationArray, "\t")
		}

		saveIsWrong := display(stream, strings.Join([]string{strings.Join(tabulationArray, ""), outputString}, ""))
		if saveIsWrong != nil { return saveIsWrong }

		if address.subnets != nil {
			err2 := SaveAddressInFile(address.subnets, level + 1, stream ,display)
			if err2 != nil { return err2 }
		}
	}

	return nil
}

func hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		if ipnet.Mask[3] == 0x00 && ip[3] != 0x00 {
			continue
		}

		ips = append(ips, ip.String())
	}

	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func arrayContainsElement(array []string, element string) bool {
	for _, s := range array {
		if s == element {
			return true
		}
	}

	return false
}
