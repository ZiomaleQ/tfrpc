package main

import (
	"encoding/json"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type Process struct {
	Name        string `json:"Name"`
	WindowTitle string `json:"mainWindowTitle"`
	StartTime   string `json:"StartTime"`
}

func main() {
	rpcErr := client.Login("622783718783844356")

	for {

		if rpcErr != nil {
			time.Sleep(2 * time.Minute)
			rpcErr = client.Login("622783718783844356")
			continue
		}

		processes := getProcessList()

		worthProcesses := make([]Process, 0)

		for _, process := range processes {
			if process.Name == "crosvm" {
				worthProcesses = append(worthProcesses, process)
			}
		}

		if len(worthProcesses) == 0 {
			time.Sleep(1 * time.Minute)
			continue
		}

		tsInt, err := strconv.ParseInt(worthProcesses[0].StartTime, 10, 64)

		if err != nil {
			panic(err)
		}

		ts := time.UnixMicro(tsInt)

		client.SetActivity(client.Activity{
			State:      "Gra z Google Play",
			Details:    "Arknights",
			LargeImage: "arknights",
			LargeText:  "Arknights",
			SmallImage: "keiko",
			SmallText:  "Keiko",
			Timestamps: &client.Timestamps{
				Start: &ts,
			},
		})

		time.Sleep(1 * time.Minute)
	}
}

func getProcessList() []Process {
	rawOut, err := runPWSHCommand("-Command", "Get-Process | Where-Object {$_.mainWindowTitle} | Select Name, mainWindowTitle, StartTime | ConvertTo-Json -Compress")

	if err != nil {
		panic(err)
	}

	rawSize, err := runPWSHCommand("-Command", "(Get-Process | Where-Object {$_.mainWindowTitle} ).Count")

	if err != nil {
		panic(err)
	}

	processList := make([]Process, int(rawSize[0]))

	json.Unmarshal([]byte(rawOut), &processList)

	for idx, process := range processList {
		processList[idx].StartTime = process.StartTime[strings.IndexRune(process.StartTime, '(')+1 : strings.IndexRune(process.StartTime, ')')]
	}

	return processList
}

func runPWSHCommand(args ...string) ([]byte, error) {
	cmd := exec.Command("powershell.exe", args...)
	return cmd.Output()
}
