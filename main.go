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

const PARAMETERS = "PARAMETERS \nfunctionmode = [send|serve|discover] \nSEND PARAMETERS \nserveraddress filename"

func main() {
	var args = os.Args[1:]
	if len(args) < 1 {
		log.Fatal( "too fiew arguments ",PARAMETERS)
	}
	switch args[0] {

  case "dump-config":

    jsonEncodedSettings,err := settings.DumpSettings()
    if err == nil{
      log.Println(jsonEncodedSettings)
    }else{
      log.Panic("unable to marshal default settings")
    }


	case "send":

	  var settings = settings.LoadSettings()

		if len(args) == 3 {

			network.Sender(args[1], settings.File_server_port, args[2],settings.Buffer_dimension)

		// if filename is not pass as argument read file from stdin
		} else if len(args) < 3 {

			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				network.Sender(args[1], settings.File_server_port, scanner.Text(),settings.Buffer_dimension)
			}
		}

		case "serve":

	  var settings = settings.LoadSettings()

		// launch discovery agent
		go discovery.Listen_discovery_requests(settings.Discovery_server_port)

		//launch server
		network.Reciver(settings.File_server_port,settings.Buffer_dimension,settings.Destination_folder)

	case "discover":

	  var settings = settings.LoadSettings()

		var available_servers = discovery.Send_discovery_request(
      settings.Discovery_client_port,
      settings.Discovery_server_port,
      settings.Broadcast_address,
      settings.Discovery_max_tries)

		for _,discoveryResponse := range available_servers{

			fmt.Printf("name: %s\n",discoveryResponse.Hostname )
			fmt.Printf("address: %s\n",discoveryResponse.Address)
		}

	}
}
