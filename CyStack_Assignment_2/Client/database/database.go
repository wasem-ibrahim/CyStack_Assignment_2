package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func SaveInDatabase(hostname string, ipAddress string, memoryUtilization string, diskUtilization string, localUsers string, runningProcesses string, installedApplications string) {

	db, err := sql.Open("sqlite3", "/Users/waseemmilhim/Documents/Coding/golang/osinfo.db")
	if err != nil {
		fmt.Println(err)

	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS os_info (hostname TEXT, ip_address TEXT, memory_utilization TEXT, disk_utilization TEXT, local_users TEXT, running_processes TEXT, installed_applications TEXT)")
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec("INSERT INTO os_info (hostname, ip_address, memory_utilization, disk_utilization, local_users, running_processes, installed_applications) VALUES (?, ?, ?, ?, ?, ?, ?)", hostname, ipAddress, memoryUtilization, diskUtilization, localUsers, runningProcesses, installedApplications)
	if err != nil {
		fmt.Println(err)
	}
}
