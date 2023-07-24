package tools

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports, results chan int, url string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", url, p)
		fmt.Printf("Scanning port: %d\n", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func TCPScanner(url string, x, y int) {
	ports := make(chan int)
	results := make(chan int)
	var openports []int

	for i := x; i <= y; i++ {
		go worker(ports, results, url)
	}

	go func() {
		for i2 := x; i2 <= y; i2++ {
			ports <- i2
		}
		close(ports)
	}()
	
	for i3 := x; i3 <= y; i3++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
func SingleTcpScanner(url string, port int) {
	address := fmt.Sprintf("%s:%d", url, port)
	conn, err := net.DialTimeout("tcp", address, 1 * time.Second)
	if err != nil {
		fmt.Printf("Error: %v ", err)
	} else {
		fmt.Print("Port is open!\n")
	}
	conn.Close()
}