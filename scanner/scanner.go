package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPorts scans a range of ports on a target host for multiple protocols.
func ScanPorts(protocols []string, target string, ports []int) map[string]map[int]string {
	results := make(map[string]map[int]string)

	for _, protocol := range protocols {
		// Initialize results map for each protocol
		if results[protocol] == nil {
			results[protocol] = make(map[int]string)
		}

		for _, port := range ports {
			status := scanPort(protocol, target, port)
			results[protocol][port] = status
		}
	}

	return results
}

// scanPort checks if a port is open.
func scanPort(protocol string, target string, port int) string {
	address := fmt.Sprintf("%s:%d", target, port)

	conn, err := net.DialTimeout(protocol, address, 2*time.Second)
	if err != nil {
		return "closed"
	}
	conn.Close()
	return "open"
}
