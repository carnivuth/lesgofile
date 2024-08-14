package main

import (
	"bufio"
	"fmt"
	"lesgofile/discovery"
	"lesgofile/network"
	"lesgofile/settings"
	"log"
	"os"
	"strconv"
)

const PARAMETERS = "PARAMETERS \nfunctionmode = [send|serve|discover] \nSEND PARAMETERS \nserveraddress filename"

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
		go discovery.Listen_discovery_requests(settings.SETTINGS["DISCOVERY_SERVER_PORT"])
		//launch server
		network.Reciver(settings.SETTINGS["PORT"])
	case "discover":
		maxTries,err:= strconv.Atoi(settings.SETTINGS["DISCOVERY_MAX_TRIES"])
		if err != nil{
		log.Panicf("unable to convert discovery parameter")
	}
		var available_servers = discovery.Send_discovery_request(settings.SETTINGS["DISCOVERY_CLIENT_PORT"],settings.SETTINGS["DISCOVERY_SERVER_PORT"],settings.SETTINGS["BROADCAST_ADDRESS"],maxTries)
		for _,discoveryResponse := range available_servers{

			fmt.Printf("name: %s\n",discoveryResponse.Hostname )
			fmt.Printf("address: %s\n",discoveryResponse.Address)
		}

	}
}
