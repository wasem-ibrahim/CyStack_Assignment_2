package main

import (
	"fmt"
	"log"
	"os"

	database "golang/database"
	networking "golang/networking"
	osinfo "golang/osinfo"
)

func main() {

	// Gather the OS information
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	ipAddress, err := osinfo.GetIPAddress()
	if err != nil {
		log.Fatal(err)
	}

	memoryUtilization, err := osinfo.GetMemoryUtilization()
	if err != nil {
		fmt.Println(err)
	}

	diskUtilization, err := osinfo.GetDiskUtilization()
	if err != nil {
		fmt.Println(err)

	}

	localUsers, err := osinfo.GetLocalUsers()
	if err != nil {
		fmt.Println(err)

	}

	runningProcesses, err := osinfo.GetRunningProcesses()
	if err != nil {
		fmt.Println(err)

	}

	installedApplications, err := osinfo.GetInstalledApplications()
	if err != nil {
		fmt.Println(err)
	}

	database.SaveInDatabase(hostname, ipAddress, memoryUtilization, diskUtilization, localUsers, runningProcesses, installedApplications)

	networking.SendData(hostname, ipAddress, memoryUtilization, diskUtilization, localUsers, runningProcesses, installedApplications)

}
