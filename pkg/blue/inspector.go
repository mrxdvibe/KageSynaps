package blue

import (
	"fmt"
	"time"
)

type TelemetryAlert struct {
	Timestamp      time.Time
	Artifact       string
	Severity       string
	DetectedBy     string
	MITRETechnique string
}

type BlueInspector struct {
	AlertsChan chan TelemetryAlert
}

func NewBlueInspector() *BlueInspector {
	return &BlueInspector{
		AlertsChan: make(chan TelemetryAlert, 50),
	}
}

func (bi *BlueInspector) AnalyzeExecution(action string, targetPID int) {
	alert := TelemetryAlert{
		Timestamp:      time.Now(),
		Artifact:       fmt.Sprintf("Process Access / Injection Attempt on PID %d via [%s]", targetPID, action),
		Severity:       "HIGH",
		DetectedBy:     "KageSynaps Telemetry Engine",
		MITRETechnique: "T1055 (Process Injection)",
	}

	bi.AlertsChan <- alert
}
