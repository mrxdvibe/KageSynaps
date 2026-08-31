package channel

import (
	"encoding/json"
	"fmt"
	"time"
)

type C2Packet struct {
	AgentID   string `json:"agent_id"`
	Command   string `json:"command,omitempty"`
	Payload   []byte `json:"payload,omitempty"`
	Timestamp int64  `json:"timestamp"`
}

type MalleableChannel struct {
	SessionKey []byte
}

func NewMalleableChannel(key []byte) *MalleableChannel {
	return &MalleableChannel{SessionKey: key}
}

func (mc *MalleableChannel) PackPayload(agentID, command string, rawData []byte) ([]byte, error) {
	packet := C2Packet{
		AgentID:   agentID,
		Command:   command,
		Payload:   rawData,
		Timestamp: time.Now().Unix(),
	}

	return json.Marshal(packet)
}

func (mc *MalleableChannel) UnpackPayload(packetBytes []byte) (*C2Packet, error) {
	var packet C2Packet
	err := json.Unmarshal(packetBytes, &packet)
	if err != nil {
		return nil, fmt.Errorf("invalid C2 frame structure: %v", err)
	}

	return &packet, nil
}
