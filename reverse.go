package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "os/exec"
    "strings"
    "time"
)

var cache map[string]struct{}

func reverseIP(ip string, recursive bool, live bool, outFile *os.File) {
    if _, ok := cache[ip]; ok {
        return
    }

    domains, err := net.LookupAddr(ip)
    if err != nil {
        return
    }
    cache[ip] = struct{}{}
    for _, domain := range domains {
        domain = strings.TrimSpace(domain)
        domain = strings.TrimRight(domain, ".")
        if domain == "" {
            continue
        }
        if live {
            fmt.Println(domain)
        }
        if outFile != nil {
            outFile.WriteString(domain + "\n")
        }
        if recursive {
            ips, err := net.LookupIP(domain)
            if err == nil {
                for _, ip := range ips {
                    reverseIP(ip.String(), true, live, outFile)
                }
            }
        }
    }
}

func recursiveSearch(ips []string, recursive bool, live bool, outFile *os.File) {
    for _, ip := range ips {
        reverseIP(ip, recursive, live, outFile)
    }
}

func main() {
    cache = make(map[string]struct{})
    banner := `
     \.\'.\     
      \'\'.\    
     __\.\:/_// ~~~~~
    {{{{{(__(")

Coded By HnvDi3
Gorev - Reverse DNS Lookup Tool

`

    recursive := false
    live := false
    args := os.Args[1:]
    if len(args) < 1 {
        fmt.Println(banner)
        fmt.Println("[!] gorev <hosts:list> <-r:recursive>\n")
        return
    }

    // Process additional arguments
    for i, arg := range args {
        if arg == "-r" || arg == "--recursive" {
            recursive = true
            args = append(args[:i], args[i+1:]...)
            break
        }
    }

    filename := args[0]
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Failed to open file:", err)
        return
    }
    defer file.Close()

    outFile, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Failed to create file:", err)
        return
    }
    defer outFile.Close()

    scanner := bufio.NewScanner(file)
    var ips []string
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        line = strings.TrimRight(line, ".")
        if net.ParseIP(line) != nil {
            ips = append(ips, line)
        } else {
            domainIPs, err := net.LookupIP(line)
            if err == nil {
                for _, ip := range domainIPs {
                    ips = append(ips, ip.String())
                }
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Failed to read file:", err)
        return
    }

    fmt.Print("[*] On progress ")
    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Print("\b/")
        time.Sleep(100 * time.Millisecond)
        fmt.Print("\b-")
        time.Sleep(100 * time.Millisecond)
        fmt.Print("\b\\")
        time.Sleep(100 * time.Millisecond)
        fmt.Print("\b|")
    }
    fmt.Print("\b")
    fmt.Println("\b | look good, be patient:)")

    recursiveSearch(ips, recursive, live, outFile)
    fmt.Println("[+] Result: finale.txt")

    mergeCmd := exec.Command("sh", "-c", fmt.Sprintf("cat %s output.txt > finale.txt", filename))
    if err := mergeCmd.Run(); err != nil {
        fmt.Println("Failed to merge file:", err)
        return
    }
}
