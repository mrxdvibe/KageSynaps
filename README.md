# KageSynaps

**Autonomous Purple-Team Engine for C2 Beaconing & Network Assessment**

[![Go](https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![Modular](https://img.shields.io/badge/Architecture-Modular-ff69b4?style=for-the-badge)](#architecture)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](#license)

[About](#about) • [Key Features](#key-features) • [Architecture](#architecture) • [Installation & Usage](#installation--usage) • [Disclaimer](#disclaimer)

---

## About

**KageSynaps** is a modular, high-performance Purple-Team testing engine written in Go. Designed to bridge the gap between offensive emulation and defensive inspection, it provides a flexible framework for C2 beaconing, malleable channels, and network telemetry analysis.

## Key Features

* **Red-Team Engine:** Custom beacon generation with customizable polling intervals and command execution.
* **Blue-Team Inspector:** Real-time network telemetry analysis and signal tracking.
* **Malleable Channels:** Encrypted and obfuscated communication layers for secure transport.
* **Zero-Hardcode Architecture:** Fully configurable runtime execution via flags and command-line interfaces.

## Installation & Usage

### Prerequisites

* Go 1.20 or higher installed on your system.

### Quick Start

1. **Clone the repository:**
   git clone https://github.com/your-username/KageSynaps.git
   
   cd KageSynaps
   
3. **Run the Node:**
   go run cmd/kage-node/main.go -host 127.0.0.1 -port 9090

4. **Global Port Forwarding (Optional):**
   If testing over public infrastructure, use any tunnel service such as `bore` or `ngrok`:
   bore local 9090 --to bore.pub
   go run cmd/kage-node/main.go -host bore.pub -port <ASSIGNED_PORT>

## Disclaimer

> **Warning:** This tool is developed strictly for educational purposes, authorized security testing, and research. Unauthorized use of this software on targets without prior mutual consent is illegal.

---

<p align="center">Developed with ❤️ by <b>MrxdVibe</b></p>
