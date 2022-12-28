package osinfo

import (
	"os/exec"
	"strings"
)

func GetLocalUsers() (string, error) {
	cmd := exec.Command("dscl", "/Local/Default", "list", "/Users")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	localUsers := strings.TrimSpace(string(output))

	return localUsers, nil
}
