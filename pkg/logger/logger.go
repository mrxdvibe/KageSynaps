package logger

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type EventLog struct {
	Timestamp string `json:"timestamp"`
	NodeID    string `json:"node_id"`
	Payload   string `json:"payload"`
}

type JSONLogger struct {
	filePath string
	mu       sync.Mutex
}

func NewJSONLogger(filePath string) *JSONLogger {
	return &JSONLogger{filePath: filePath}
}

func (l *JSONLogger) WriteLog(nodeID, payload string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	entry := EventLog{
		Timestamp: time.Now().Format(time.RFC3339),
		NodeID:    nodeID,
		Payload:   payload,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(l.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(string(data) + "\n"); err != nil {
		return err
	}

	return nil
}
