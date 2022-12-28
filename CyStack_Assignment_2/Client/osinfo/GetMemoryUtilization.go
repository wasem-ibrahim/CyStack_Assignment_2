package osinfo

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetMemoryUtilization() (string, error) {
	output, err := exec.Command("top", "-l", "1", "-n", "0").Output()
	if err != nil {
		return "", err
	}

	utilizationStr := strings.TrimSpace(string(output))
	utilizationStr = strings.Split(utilizationStr, "PhysMem:")[1]

	fields := strings.Fields(utilizationStr)
	usedMemoryStr := fields[0]
	totalMemoryStr := fields[2]

	usedMemory, err := strconv.ParseFloat(strings.TrimSuffix(usedMemoryStr, "G"), 64)
	if err != nil {
		return "", err
	}

	totalMemory, err := strconv.ParseFloat(strings.TrimPrefix(strings.TrimSuffix(totalMemoryStr, "M"), "("), 64)
	if err != nil {

		return "", err
	}

	usedMemory /= 1024
	totalMemory /= 1024

	utilization := (usedMemory / totalMemory) * 100

	return fmt.Sprintf("%.2f%%", utilization), nil
}
