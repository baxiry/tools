package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	host      string
	startPort int
	endPort   int
	wg        sync.WaitGroup
)

func main() {
	userInput()
}

func userInput() {
	fmt.Println("Type Host> ")
	fmt.Scan(&host)
	fmt.Println("Type starting Port (i.e. 80)> ")
	fmt.Scan(&startPort)
	fmt.Println("Type end Port (i.e. 9999)> ")
	fmt.Scan(&endPort)
	fmt.Println("Runing Scan... ")

	portRange := endPort - startPort
	endPortSet1 := (portRange / 10) + startPort
	endPortSet2 := (portRange / 10) + endPortSet1
	endPortSet3 := (portRange / 10) + endPortSet2
	endPortSet4 := (portRange / 10) + endPortSet3
	endPortSet5 := (portRange / 10) + endPortSet4
	endPortSet6 := (portRange / 10) + endPortSet5
	endPortSet7 := (portRange / 10) + endPortSet6
	endPortSet8 := (portRange / 10) + endPortSet7
	endPortSet9 := (portRange / 10) + endPortSet8

	wg.Add(10) // 1min to run 65 ports on 10 concurrent goup
	go checkPort(host, startPort, endPortSet2)
	go checkPort(host, (endPortSet1 + 1), endPortSet2)
	go checkPort(host, (endPortSet2 + 1), endPortSet3)
	go checkPort(host, (endPortSet3 + 1), endPortSet4)
	go checkPort(host, (endPortSet4 + 1), endPortSet5)
	go checkPort(host, (endPortSet5 + 1), endPortSet6)
	go checkPort(host, (endPortSet6 + 1), endPortSet7)
	go checkPort(host, (endPortSet7 + 1), endPortSet8)
	go checkPort(host, (endPortSet8 + 1), endPortSet9)
	go checkPort(host, (endPortSet9 + 1), endPort)
	wg.Wait()
}

func checkPort(host string, startPort, endPort int) {
	for i := startPort; i <= endPort; i++ {
		qualifiedHost := fmt.Sprintf("%s%s%d", host, ":", i)
		conn, err := net.DialTimeout("tcp", qualifiedHost, 10*time.Millisecond)
		if err != nil {
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n1\n22\n\n\n\n")
		conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))

		//buff := make([]byte, 1024)
		//n, _ := conn.Read(buff)
		//fmt.Printf("Port: %d%s\n", i, buff[:n])
		fmt.Println("Port: ", i)
	}
	wg.Done()
}
