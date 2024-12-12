package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cindyhy98/go-portscanner/scanner"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask the user for the target
	fmt.Print("Enter the target to scan (e.g., google.com): ")
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	protocols := []string{"tcp", "udp"}
	ports := []int{80, 443, 53, 22}

	fmt.Println("Starting port scan...")

	// Check the scan starting time
	start := time.Now()

	results := scanner.ScanPorts(protocols, target, ports)

	for protocol, portResults := range results {
		fmt.Printf("Protocol: %s\n", protocol)
		for port, status := range portResults {
			fmt.Printf("Port %d: %s\n", port, status)
		}
	}

	// Check the scan duration
	duration := time.Since(start)
	fmt.Printf("Scan completed in %s\n", duration)
}
