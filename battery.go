package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func battery() {
	percentRegex := regexp.MustCompile("\\d{1,3}%")
	statusRegex := regexp.MustCompile("Discharging|Charging")

	out, err := exec.Command("acpi").Output()
	if err != nil {
		fmt.Print(err)
	}

	battery := percentRegex.FindString(string(out))
	status := statusRegex.FindString(string(out))
	value, err := strconv.Atoi(strings.TrimSuffix(battery, "%"))
	if err != nil {
		fmt.Print(err)
	}

	switch {
	case value < 5:
		battery = "<span foreground='#F50000'> " + battery + "</span>"
	case value < 25:
		battery = "<span foreground='#F8FF36'> " + battery + "</span>"
	case value < 50:
		battery = " " + battery
	case value < 75:
		battery = " " + battery
	default:
		battery = " " + battery
	}

	switch status {
	case "Charging":
		battery = battery + " "
	case "Discharging":
		battery = battery + " "
	}

	fmt.Print(battery)
}
