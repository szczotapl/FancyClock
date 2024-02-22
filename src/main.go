package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var HOME string
var fontDir string

const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

func init() {
	var err error
	HOME, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	fontDir = filepath.Join(HOME, ".gitman/packages/fancyclock/src")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error clearing screen:", err)
	}
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

		cmd := exec.Command("toilet", "-d", fontDir, "-f", "font", "--filter", "metal", fmt.Sprintf("%s%s%s", color, asciiTime, colorReset))
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			fmt.Println("Error executing toilet command:", err)
		}

		time.Sleep(1 * time.Second)
	}
}
