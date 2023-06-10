package main

import (
	"fmt"
	"os"

	"github.com/carnivuth/lesgofile/network"
	"github.com/carnivuth/lesgofile/settings"
)

// /check parameters and launch slaves
func main() {
	var terminate chan int = make(chan int)
	settings.LoadSettings()
	if len(os.Args) < 2 {
		fmt.Println("too fiew arguments ")
		return
	}
	args := os.Args[1:]
	if args[0] == "send" {

		network.Sender(args[1], settings.SETTINGS["PORT"], args[2])

	} else if args[0] == "recive" {

		//launch server
		network.Reciver(settings.SETTINGS["PORT"], terminate)

	}
}
