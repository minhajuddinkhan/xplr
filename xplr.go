package main

import (
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/minhajuddinkhan/xplr/cmd"
)

func main() {

	t1 := time.Now()
	err := cmd.Root().Execute()
	t2 := t1.Sub(time.Now())

	if err != nil {
		color.New(color.FgRed).Println("Oops!", err)
		os.Exit(1)
	} else {
		c := color.New(color.FgWhite).Add(color.Bold)
		c.Println("\nExecution time:: ", t2)
	}
}
