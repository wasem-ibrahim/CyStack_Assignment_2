package osinfo

import (
	"bufio"
	"os/exec"
	"strings"
)

func GetIPAddress() (string, error) {
	cmd := exec.Command("ifconfig")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	var ipAddress string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "inet ") {
			fields := strings.Fields(line)
			for _, field := range fields {
				if strings.Contains(field, ".") {
					ipAddress = field
					break
				}
			}
		}
		if ipAddress != "" {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return ipAddress, nil
}
