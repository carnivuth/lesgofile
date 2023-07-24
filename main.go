package main

import (
	"bufio"
	"os"

	"github.com/carnivuth/lesgofile/logger"
	"github.com/carnivuth/lesgofile/network"
	"github.com/carnivuth/lesgofile/settings"
)

// /check parameters and launch slaves

const PARAMETERS = "PARAMETERS \nfunctionmode = [send|recive] \nSEND PARAMETERS \nserveraddress filename"

// PARAMETERS

// functionmode = [send|recive]
// SEND PARAMETERS
//
//	serveraddress filename
func printToStdOut(message string) {
	println(message)
}

func main() {
	logger.AddListener(logger.Log, printToStdOut)
	settings.LoadSettings()

	logger.Emit(logger.Log, "settings loaded")

	var args = os.Args[1:]
	if len(args) < 1 {
		logger.Emit(logger.Log, "too fiew arguments ")
		logger.Emit(logger.Log, PARAMETERS)
		return
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

	} else if args[0] == "recive" {

		//launch server
		network.Reciver(settings.SETTINGS["PORT"])

	}
}
