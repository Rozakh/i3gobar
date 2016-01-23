package main

import (
	"fmt"
	"os/exec"
)

func power() {
	if clickedButton == "1" {
		echoCmd := exec.Command("echo", "-e", " Logout\n Reboot\n Poweroff")
		rofiCmd := exec.Command("rofi", "-dmenu", "-i", "-p", "Shutdown Menu:", "-bg", "#222222", "-fg", "#888888",
			"-hlbg", "#285577", "-hlfg", "#ffffff", "-width", "20", "-hide-scrollbar")
		echoOut, err := echoCmd.StdoutPipe()
		if err != nil {
			fmt.Print(err)
		}
		rofiCmd.Stdin = echoOut
		echoCmd.Start()
		result, _ := rofiCmd.Output()
		switch string(result) {
		case " Logout\n":
			exec.Command("i3-msg", "exit").Run()
		case " Reboot\n":
			exec.Command("reboot").Run()
		case " Poweroff\n":
			exec.Command("poweroff").Run()
		}
	}
}
