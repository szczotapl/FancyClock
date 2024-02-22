package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const customFontURL = "https://raw.githubusercontent.com/riviox/FancyClock/master/src/assets/font.flf"
const customFontFile = "/tmp/font.flf"

// ANSI escape codes for colors
const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

func downloadFont(url, filePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	return err
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printColoredAsciiClock() {
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

	cmd := exec.Command("toilet", "-f", customFontFile, "--filter", "metal", fmt.Sprintf("%s%s%s", color, asciiTime, colorReset))
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	err := downloadFont(customFontURL, customFontFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for {
		printColoredAsciiClock()
		time.Sleep(1 * time.Second)
	}
}
