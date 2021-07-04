package ipAddress

import (
	"errors"
	"strings"
)

func repeatedAddElements(stringArray []*string, actualString []string, index *int, master *IPAddress, level int) ([]string, *IPAddress,  error) {
	slave, _ := CreateNewIpAddress(strings.Replace(actualString[0], "\t", "", level), actualString[1])

	if master != nil {
		master.subnets = append(master.subnets, slave)
	}

	*index++

	if *index >= len(stringArray) {
		return nil, nil, errors.New("out of range")
	}

	return strings.Split(*stringArray[*index], ","), slave, nil
}

func createNextLevel(actualString []string, index *int, stringArray []*string, master *IPAddress, level int) error {
	var addElementError error

	if strings.Count(actualString[0], "\t") == 0 {
		return errors.New("level end")
	}

	if strings.Count(actualString[0], "\t") == level {
		for *index <= len(stringArray) - 1 {
			var slave *IPAddress
			actualString, slave, addElementError = repeatedAddElements(stringArray, actualString, index, master, level)
			if addElementError != nil {
				return errors.New("end")
			}

			err := createNextLevel(actualString, index, stringArray, slave, level + 1)

			if err != nil {
				return errors.New("level end")
			}
		}
	}

	return nil
}

func checkRepeat(ipAddresses []*IPAddress) []*IPAddress {
	var correctAddressList []*IPAddress
	var subnetsCorrect []*IPAddress

	for x, address := range ipAddresses {
		repeat := false

		if address.subnets != nil {
			subnetsCorrect = checkRepeat(address.subnets)
		}

		for y, ipAddress := range ipAddresses {
			if x == y {
				break
			}

			if address.getAddressString() == ipAddress.getAddressString() &&
				address.getAddressMaskString() == ipAddress.getAddressMaskString() {
				repeat = true
			}
		}

		if !repeat {
			master := ipAddresses[x]

			master.subnets = subnetsCorrect
			correctAddressList = append(correctAddressList, master)
		} else {
			println("\033[31mRepeated address", ipAddresses[x].getAddressString(), "\033[0m")
		}
	}

	return correctAddressList
}

func AddressesMapper(stringArray []*string) []*IPAddress {
	var ipAddresses []*IPAddress
	var addElementError error

	index := 0

	for index <= len(stringArray) - 1 {
		actualString := strings.Split(*stringArray[index], ",")

		var master *IPAddress
		actualString, master, addElementError = repeatedAddElements(stringArray, actualString, &index, nil, 0)
		if addElementError != nil {
			break
		}

		err := createNextLevel(actualString, &index, stringArray, master, 1)

		if err != nil {
			ipAddresses = append(ipAddresses, master)
		}
	}

	return checkRepeat(ipAddresses)
}
