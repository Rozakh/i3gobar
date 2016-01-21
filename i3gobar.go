package main

import (
	"flag"
	"os"
)

var clickedButton string

func main() {
	clickedButton = os.Getenv("BLOCK_BUTTON")
	volumeFlag := flag.Bool("v", false, "Show volume info")
	batteryFlag := flag.Bool("b", false, "Show battery info")
	mailFlag := flag.Bool("m", false, "Show mail info")
	powerFlag := flag.Bool("p", false, "Show power menu")
	flag.Parse()

	if *volumeFlag {
		volume()
	}

	if *batteryFlag {
		battery()
	}

	if *mailFlag {
		mail()
	}

	if *powerFlag {
		power()
	}
}
