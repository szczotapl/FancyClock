package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	font        = "~/.gitman/packages/FancyClock/assets/ansireg.flf"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		clearScreen()
		currentTime := time.Now()
		asciiTime := currentTime.Format("15:04:05")

		color := colorReset
		second := currentTime.Second()
		switch {
		case second%2 == 0:
			color = colorGreen
		default:
			color = colorYellow
		}

		cmd := exec.Command("toilet", "-f", font, "--filter", "metal", fmt.Sprintf("%s%s%s", color, asciiTime, colorReset))
		cmd.Stdout = os.Stdout
		cmd.Run()
		time.Sleep(1 * time.Second)
	}
}
