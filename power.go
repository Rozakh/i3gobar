package main

import "os/exec"

func power() {
	switch clickedButton {
	case "1":
		err := exec.Command("zenity", "--question", "--text=Shutdown. Are you sure?").Run()
		if err == nil {
			exec.Command("poweroff").Run()
		}
	case "3":
		err := exec.Command("zenity", "--question", "--text=Restart. Are you sure?").Run()
		if err == nil {
			exec.Command("reboot").Run()
		}
	}
}
