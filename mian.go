package main

import (
	"fmt"
	"github.com/apparentlymart/go-cidr/cidr"
	"log"
	"net"
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

func isV4IP(ip net.IP) bool {
	return ip.To4() != nil
}

func parseCIDR(s string) (net.IP, net.IP) {
	ip, network, err := net.ParseCIDR(s)
	if err != nil {
		return nil, nil
	}

	start, end := cidr.AddressRange(network)
	prefixLen, _ := network.Mask.Size()

	if isV4IP(ip) && prefixLen < 31 {
		start = cidr.Inc(start)
		end = cidr.Dec(end)
	}

	return start, end
}

func Hosts(cidr string) ([]string, error) {
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

//  http://play.golang.org/p/m8TNTtygK0
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

func main() {
	argsWithoutProg := os.Args[1:]

	for i, s := range argsWithoutProg {
		switch s {
		case "--file":
			mainAddressesArray := fileOption(argsWithoutProg, i)

			for _, address := range mainAddressesArray {
				onesMask, _ := address.IpAddress.Mask.Size()

				masterSubnets, _ := Hosts(fmt.Sprintf("%s/%d", address.IpAddress.IP.String(), onesMask))

				var subnetsAddresses [][]string

				for _, subnet := range address.Subnets {
					onesMask, _ := subnet.IpAddress.Mask.Size()

					subnetSubnets, _ := Hosts(fmt.Sprintf("%s/%d", subnet.IpAddress.IP.String(), onesMask))

					subnetsAddresses = append(subnetsAddresses, subnetSubnets)
				}

				var repeatList []string

				for _, subnet := range masterSubnets {
					for _, subnetsAddress := range subnetsAddresses {
						for _, s2 := range subnetsAddress {
							if subnet == s2 {
								repeatList = append(repeatList, subnet)
							}
						}
					}
				}

				var correctAddressList []string

				for _, subnet := range masterSubnets {
					if !arrayContainsElement(repeatList, subnet) {
						correctAddressList = append(correctAddressList, subnet)
					}
				}

				for _, s2 := range correctAddressList {
					addressString := fmt.Sprintf("%s/%d", s2, 24)
					slave, _ := ipAddress.CreateNewIpAddress(addressString, "WOLNA")

					address.Subnets = append(address.Subnets, slave)
				}

				//start, end := parseCIDR(fmt.Sprintf("%s/%d", address.IpAddress.IP.String(), onesMask))

				//println(start.String(), end.String())
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
