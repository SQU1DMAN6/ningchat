package main

import (
	"fmt"
	"os"

	boot "github.com/SQU1DMAN6/ningchat/app"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must provide an argument to start the NingChat server.")
		os.Exit(1)
	}

	arg1 := os.Args[1]
	if arg1 == "production" {
		boot.BootApp(true)
	} else if arg1 == "local" {
		boot.BootApp(false)
	} else {
		fmt.Println("You must provide either 'production' if running on the InkDrop machine, or 'local' if running on your local device.")
		os.Exit(1)
	}
}
