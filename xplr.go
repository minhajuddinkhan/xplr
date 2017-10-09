package main

import (
	"fmt"
	"os"
	"time"

	"github.com/minhajuddinkhan/xplr/cmd"
)

func main() {

	t1 := time.Now()
	err := cmd.Root().Execute()
	t2 := t1.Sub(time.Now())

	if err != nil {
		fmt.Errorf("error:", err)
		os.Exit(1)
	} else {
		fmt.Println("Execution time/ ", t2)
	}
}
