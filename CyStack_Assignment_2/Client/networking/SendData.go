package networking

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
)

func SendData(hostname string, ipAddress string, memoryUtilization string, diskUtilization string, localUsers string, runningProcesses string, installedApplications string) {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	data := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n", hostname, ipAddress, memoryUtilization, diskUtilization, localUsers, runningProcesses, installedApplications)
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Println(err)
		return
	}

	client := http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return conn, nil
			},
		},
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: request failed with status code", res.StatusCode)
		return
	}

	fmt.Println(string(body))
}
