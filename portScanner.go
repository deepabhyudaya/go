package main

import (
    "fmt"
    "net"
    "sync"
    "time"
    "os"
)

func scanPort(wg *sync.WaitGroup, host string, port int) {
    defer wg.Done()
    address := fmt.Sprintf("%s:%d", host, port)
    conn, err := net.DialTimeout("tcp", address, time.Second)
    if err == nil {
        fmt.Printf("Port %d is open\n", port)
        conn.Close()
    }
}

func main() {
    target := os.Getenv("HOST")
    startPort := os.Getenv("SPORT")
    endPort := os.Getenv("EPORT")

    fmt.Printf("Scanning ports %d to %d on host %s\n", startPort, endPort, target)

    var wg sync.WaitGroup
    for port := startPort; port <= endPort; port++ {
        wg.Add(1)
        go scanPort(&wg, target, port)
    }

    wg.Wait()
    fmt.Println("Scan complete!")
}
