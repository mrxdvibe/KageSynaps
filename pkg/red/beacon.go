package red

import (
	"fmt"
	"time"
)

type RedBeacon struct {
	NodeID    string
	IsStealth bool
}

func NewRedBeacon(id string) *RedBeacon {
	return &RedBeacon{
		NodeID:    id,
		IsStealth: true,
	}
}

func (rb *RedBeacon) SimulateAttackTask(taskName string) {
	fmt.Printf("[RED] Executing Attack Simulation: %s on Node [%s]\n", taskName, rb.NodeID)
	time.Sleep(500 * time.Millisecond)
}
