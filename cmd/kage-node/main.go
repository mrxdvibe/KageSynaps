package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"KageSynaps/pkg/blue"
	"KageSynaps/pkg/channel"
	"KageSynaps/pkg/red"
)

func main() {
	fmt.Println("==============================================================")
	fmt.Println("  KAGESYNAPS v1.0.0 - Autonomous Purple-Team Engine           ")
	fmt.Println("  Developer: MrxdVibe                                         ")
	fmt.Println("==============================================================")

	beacon := red.NewRedBeacon("KS-NODE-ALPHA")
	inspector := blue.NewBlueInspector()
	dummyKey := []byte("0123456789abcdef0123456789abcdef")
	malleable := channel.NewMalleableChannel(dummyKey)

	// Polling Loop: Polling Server for tasks & Sending Telemetry
	go func() {
		for {
			time.Sleep(3 * time.Second)

			conn, err := net.Dial("tcp", "127.0.0.1:9090")
			if err != nil {
				continue
			}

			packedData, _ := malleable.PackPayload(beacon.NodeID, "", nil)
			conn.Write(packedData)

			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err == nil && n > 0 {
				respPacket, err := malleable.UnpackPayload(buf[:n])
				if err == nil && respPacket.Command != "" {
					fmt.Printf("\n[*] Received Command from C2: %s\n", respPacket.Command)
					beacon.SimulateAttackTask(respPacket.Command)
					inspector.AnalyzeExecution(respPacket.Command, 1337)
				}
			}
			conn.Close()
		}
	}()

	// Alert Telemetry Listener
	go func() {
		for alert := range inspector.AlertsChan {
			alertMsg := fmt.Sprintf("[%s] %s | Severity: %s | Technique: %s",
				alert.Timestamp.Format("15:04:05"),
				alert.Artifact,
				alert.Severity,
				alert.MITRETechnique,
			)

			conn, err := net.Dial("tcp", "127.0.0.1:9090")
			if err != nil {
				continue
			}

			packedData, _ := malleable.PackPayload(beacon.NodeID, "", []byte(alertMsg))
			conn.Write(packedData)
			conn.Close()
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\n[*] KageSynaps Node shutting down safely...")
}
