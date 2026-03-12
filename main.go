package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: portcheck <host:port> [timeout_ms]")
		fmt.Println("Example: portcheck localhost:8080 2000")
		os.Exit(1)
	}

	addr := os.Args[1]
	timeout := 3 * time.Second

	if len(os.Args) >= 3 {
		ms, err := time.ParseDuration(os.Args[2] + "ms")
		if err == nil {
			timeout = ms
		}
	}

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		fmt.Printf("❌ %s is not reachable (%v)\n", addr, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("✅ %s is open\n", addr)
}
