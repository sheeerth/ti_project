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

func main() {
	argsWithoutProg := os.Args[1:]

	for i, s := range argsWithoutProg {
		switch s {
		case "--file":
			mainAddressesArray := fileOption(argsWithoutProg, i)

			for _, address := range mainAddressesArray {
				// _, ipv4Net, _ := net.ParseCIDR("157.158.1.0/24")
				// b := address.IpAddress.Contains(ipv4Net.IP)

				onesMask, _ := address.IpAddress.Mask.Size()

				start, end := parseCIDR(fmt.Sprintf("%s/%d", address.IpAddress.IP.String(), onesMask))

				//println(strings.Split(parseRange.String(), "-")[0])
				println(start.String(), end.String())
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
