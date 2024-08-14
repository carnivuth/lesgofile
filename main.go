package main

import (
	"bufio"
	"fmt"
	"lesgofile/discovery"
	"lesgofile/network"
	"lesgofile/settings"
	"log"
	"os"
)

const PARAMETERS = "PARAMETERS \nfunctionmode = [send|serve] \nSEND PARAMETERS \nserveraddress filename"

func main() {
	settings.LoadSettings()
	var args = os.Args[1:]
	if len(args) < 1 {
		log.Fatal( "too fiew arguments ",PARAMETERS)
	}
	switch args[0] {
	case "send":
		if len(args) == 3 {
			network.Sender(args[1], settings.SETTINGS["PORT"], args[2])
		} else if len(args) < 3 {
			//filter behavior
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				network.Sender(args[1], settings.SETTINGS["PORT"], scanner.Text())
			}
		}
		case "serve":
		// launch discovery agent
		go discovery.Listen_discovery_requests()
		//launch server
		network.Reciver(settings.SETTINGS["PORT"])
	case "discover":
		var available_servers = discovery.Send_discovery_request()
		for _,discoveryResponse := range available_servers{

			fmt.Printf("name: %s\n",discoveryResponse.Hostname )
			fmt.Printf("address: %s\n",discoveryResponse.Address)
		}

	}
}
