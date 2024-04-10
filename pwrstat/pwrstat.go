// Package pwrstat provides an interface to query the status of a UPS (Uninterruptible Power Supply)
// using the `pwrstat` command available on certain systems. It supports generating the status output
// both as a structured Go type and as JSON.
package pwrstat

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strings"
)

// Pwrstat represents the status of a UPS, including raw content and a structured status.
type Pwrstat struct {
	Content string            `json:"-"`
	Status  map[string]string `json:"status"`
}

// NewFromSystem queries the UPS status using the `pwrstat` system command.
// It returns a Pwrstat instance or an error if the command fails or is not found.
func NewFromSystem() (*Pwrstat, error) {
	output, err := exec.Command("pwrstat", "-status").Output()
	if err != nil {
		return nil, errors.New("pwrstat command failed or not found")
	}

	return parseOutput(string(output)), nil
}

// NewFromFile reads the UPS status from a specified file path.
// This is useful for testing or reading from a pre-saved status file.
func NewFromFile(path string) (*Pwrstat, error) {
	output, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parseOutput(string(output)), nil
}

// JSON serializes the Pwrstat instance into a JSON string.
func (p *Pwrstat) JSON() string {
	jsonOutput, _ := json.Marshal(p)
	return string(jsonOutput)
}

// String returns the JSON representation of the Pwrstat instance.
// It is a convenience wrapper around the JSON method.
func (p *Pwrstat) String() string {
	return p.JSON()
}

// parseOutput constructs a Pwrstat instance from the raw output string.
// It parses the output line by line and fills the Status map.
func parseOutput(content string) *Pwrstat {
	status := make(map[string]string)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if trimmedLine := strings.TrimSpace(line); len(trimmedLine) > 0 {
			parts := strings.Split(trimmedLine, ". ")
			for i := 0; i < len(parts)-1; i += 2 {
				key := strings.ReplaceAll(parts[i], ".", "")
				value := parts[i+1]
				status[key] = value
			}
		}
	}

	return &Pwrstat{
		Content: content,
		Status:  status,
	}
}
