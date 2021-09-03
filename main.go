package main

import (
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {

	hawkData, err := readJson()
	if err != nil {
		color.Red("[ERROR] main.readJson Error Executing Func: %s\n%s", err, "Closing Program in 3 Seconds...")
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}

	writeErr := writeHawk(*&hawkData)
	if writeErr != nil {
		color.Red("[ERROR] main.WriteHawk Error In Function: %s\n%s", err, "Closing Program in 3 Seconds...")
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}
	color.HiCyan("[COMPLETE] All Operations Complete. Closing Program in 3 seconds...")
	time.Sleep(time.Second * 3)
	os.Exit(0)
}
