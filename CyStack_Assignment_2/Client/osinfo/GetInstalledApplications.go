package osinfo

import (
	"os/exec"
	"strings"
)

func GetInstalledApplications() (string, error) {
	cmd := exec.Command("ls", "/Applications")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	installedApplications := strings.TrimSpace(string(output))

	return installedApplications, nil
}
