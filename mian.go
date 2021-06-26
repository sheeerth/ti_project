package main

import (
    "bufio"
    "fmt"
    "os"
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

        extractedAdress := strings.ReplaceAll(splitString[0], " ", "")

        adress := IPAddress{ipAddress: extractedAdress, description: splitString[1], level: 0}

        ipAddresses = append(ipAddresses, adress)
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