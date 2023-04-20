package dasung

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func FindDasungI2CDevicePaths() ([]string, error) {
	cmd := exec.Command("ddcutil", "detect", "--verbose")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return nil, err
	}

	output := out.String()
	lines := strings.Split(output, "\n")
	currentPath := ""
	devicePaths := make([]string, 0)

	for _, line := range lines {
		if strings.Contains(line, "I2C bus:") {
			currentPath = strings.TrimSpace(strings.Split(line, ":")[1])
		}

		if currentPath != "" && (strings.Contains(line, "DSC") || strings.Contains(line, "Dasung") || strings.Contains(line, "Paperlike")) {
			devicePaths = append(devicePaths, currentPath)
			currentPath = ""
		}
	}

	if len(devicePaths) == 0 {
		return nil, errors.New("no Dasung Paperlike displays found")
	}
	return devicePaths, nil
}
