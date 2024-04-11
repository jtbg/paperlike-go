package dasung

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func FindDasungI2CDevicePaths() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "ddcutil", "detect", "--verbose")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return nil, fmt.Errorf("failed to run ddcutil: %v", err)
	}

	ddcOutput := out.String()
	ddcOutputLines := strings.Split(ddcOutput, "\n")
	currentPath := ""
	devicePaths := make([]string, 0)

	for _, line := range ddcOutputLines {
		if strings.Contains(line, "I2C bus:") {
			currentPath = strings.TrimSpace(strings.Split(line, ":")[1])
		}

		if currentPath != "" && (strings.Contains(line, "DSC") || strings.Contains(line, "Dasung") || strings.Contains(line, "Paperlike")) {
			log.Println(line, "\n found at ", currentPath)
			devicePaths = append(devicePaths, currentPath)
			currentPath = ""
		}
	}

	if len(devicePaths) == 0 {
		return nil, errors.New("no Dasung Paperlike displays found")
	}
	return devicePaths, nil
}
