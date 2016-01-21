package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func volume() {
	switch clickedButton {
	case "3":
		exec.Command("amixer", "-q", "set", "Master", "toggle").Run()
	case "4":
		exec.Command("amixer", "-q", "set", "Master", "2%+", "unmute").Run()
	case "5":
		exec.Command("amixer", "-q", "set", "Master", "2%-", "unmute").Run()
	}

	valueRegex := regexp.MustCompile("\\d{1,3}%")
	muteRegex := regexp.MustCompile("\\[on]|\\[off]")

	out, err := exec.Command("amixer", "get", "Master").Output()
	if err != nil {
		fmt.Print(err)
	}

	volume := valueRegex.FindString(string(out))
	mute := muteRegex.FindString(string(out))
	value, err := strconv.Atoi(strings.TrimSuffix(volume, "%"))
	if err != nil {
		fmt.Print(err)
	}

	switch {
	case value == 0:
		volume = " " + volume
	case value < 50:
		volume = " " + volume
	default:
		volume = " " + volume
	}

	if mute == "[off]" {
		volume = "<span foreground='#F8FF36'>" + volume + "</span>"
	}

	fmt.Print(volume)
}
