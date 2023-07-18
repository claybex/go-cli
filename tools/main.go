package tools

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
)
func TcpScanner(address, port string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan string, 1)
// Instead of go routine and 
	go func() {
		conn, err := net.Dial("tcp", address + ":" + port)
		if err != nil {
			log.Fatalf("Fatal error: %s \n", err)
		}
		defer conn.Close()

		ch <- "Port is open!"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("Connection timeout")
	}
}