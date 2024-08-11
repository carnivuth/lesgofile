package main

import (
	"bufio"
	"os"
	"log"
	"lesgofile/network"
	"lesgofile/settings"
)

const PARAMETERS = "PARAMETERS \nfunctionmode = [send|serve] \nSEND PARAMETERS \nserveraddress filename"

func main() {
	settings.LoadSettings()


	var args = os.Args[1:]
	if len(args) < 1 {
		log.Fatal( "too fiew arguments ",PARAMETERS)
	}
	if args[0] == "send" {
		if len(args) == 3 {
			network.Sender(args[1], settings.SETTINGS["PORT"], args[2])
		} else if len(args) < 3 {
			//filter behavior
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				network.Sender(args[1], settings.SETTINGS["PORT"], scanner.Text())

			}

		}

	} else if args[0] == "serve" {

		//launch server
		network.Reciver(settings.SETTINGS["PORT"])

	}
}
