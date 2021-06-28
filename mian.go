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

        ipv4Addr, ipv4Net, err := net.ParseCIDR(strings.ReplaceAll(splitString[0], " ", ""))
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("4-byte representation : ", ipv4Addr.To4())

        address := IPAddress{ipAddress: *ipv4Net, description: splitString[1]}

        ipAddresses = append(ipAddresses, address)
    }

    fmt.Println(ipAddresses)

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