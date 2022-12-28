package osinfo

import (
	"strings"

	"github.com/shirou/gopsutil/process"
)

func GetRunningProcesses() (string, error) {

	processes, err := process.Processes()
	if err != nil {
		return "", err
	}

	names := []string{}
	for _, proc := range processes {
		name, err := proc.Name()
		if err == nil {
			names = append(names, name)
		}
	}

	runningProcesses := strings.Join(names, "\n")

	return runningProcesses, nil
}
