package main

import (
	"flag"
	"fmt"
)

func main() {
	// İstifadəçi terminaldan -host və -port verə bilsin
	host := flag.String("host", "127.0.0.1", "C2 Server Host / Tunnel Address")
	port := flag.Int("port", 9090, "C2 Server Port")
	flag.Parse()

	serverAddr := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Printf("[+] Connecting to C2 Server at: %s\n", serverAddr)

	// Bağlantı məntiqini serverAddr vasitəsilə işə sal
}
