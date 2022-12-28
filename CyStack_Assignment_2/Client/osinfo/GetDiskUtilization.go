package osinfo

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetDiskUtilization() (string, error) {
	output, err := exec.Command("df", "-h").Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	utilizationStr := lines[1]

	fields := strings.Fields(utilizationStr)
	usedMemoryStr := fields[3]
	totalMemoryStr := fields[1]

	totalMemory, err := strconv.ParseFloat(strings.TrimSuffix(totalMemoryStr, "Gi"), 64)
	if err != nil {
		return "", err
	}

	usedMemory, err := strconv.ParseFloat(strings.TrimSuffix(usedMemoryStr, "Gi"), 64)
	if err != nil {
		return "", err
	}

	usedMemory = totalMemory - usedMemory

	utilization := (usedMemory / totalMemory) * 100

	return fmt.Sprintf("%.2f%%", utilization), nil
}
