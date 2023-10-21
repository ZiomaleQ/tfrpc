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
	err := client.Login("622783718783844356")
	if err != nil {
		panic(err)
	}

	for {
		processes := getProcessList()

		worthProcesses := make([]Process, 0)

		for _, process := range processes {
			if process.Name == "crosvm" {
				worthProcesses = append(worthProcesses, process)
			}
		}

		tsInt, err := strconv.ParseInt(worthProcesses[0].StartTime, 10, 64)

		if err != nil {
			panic(err)
		}

		ts := time.Unix(tsInt, 0)

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

		time.Sleep(15 * time.Second)
	}
}

func getProcessList() []Process {
	size := exec.Command("powershell.exe", "-Command", "(Get-Process | Where-Object {$_.mainWindowTitle} ).Count")
	cmd := exec.Command("powershell.exe", "-Command", "Get-Process | Where-Object {$_.mainWindowTitle} | Select Name, mainWindowTitle, StartTime | ConvertTo-Json -Compress")

	rawOut, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	rawSize, err := size.Output()

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
