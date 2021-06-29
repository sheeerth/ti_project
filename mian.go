package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    _ "net"
    "os"
    _ "strconv"
    "strings"
)

func openFile(filePath string) {
    var ipAddresses []IPAddress

    fmt.Println(filePath)

    f, err := os.Open(filePath)

    if err != nil {
        fmt.Println(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {

        line := scanner.Text()

        splitString := strings.Split(line, ",")

        _, ipv4Net, err := net.ParseCIDR(strings.ReplaceAll(splitString[0], " ", ""))
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("4-byte representation : ", ipv4Net.IP.To4(), ipv4Net.Mask.String())

        address := IPAddress{ipAddress: *ipv4Net, description: splitString[1]}

        ipAddresses = append(ipAddresses, address)
    }

    var correctAddressList []IPAddress

    for x, address := range ipAddresses {
        repeat := false

        for y, ipAddress := range ipAddresses {
            if x == y {
                break
            }

            if address.ipAddress.IP.String() == ipAddress.ipAddress.IP.String() && address.ipAddress.Mask.String() == ipAddress.ipAddress.Mask.String() {
                repeat = true
            }
        }

        if !repeat {
            correctAddressList = append(correctAddressList, ipAddresses[x])
        }
    }

    fmt.Println("----")

    for _, address := range correctAddressList {
        fmt.Println("4-byte representation : ", address.ipAddress.IP.To4(), address.ipAddress.Mask.String(), address.description)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func main() {
    argsWithoutProg := os.Args[1:]

    for i, s := range argsWithoutProg {
        switch s {
        case "--file":
            openFile(argsWithoutProg[i+1])
            break
        }
    }
}