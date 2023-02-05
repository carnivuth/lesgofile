package lesGoFile

import (
	"fmt"
	"net"
	"os"
)

// /check parameters and launch slaves
func main() {
	loadSettings()
	args := os.Args[1:]

	if args[0] == "send" {

		//check parameters and launch client
		if checkParameters(args[1]) {

			go client(args[1], SETTINGS["port"])

		} else {

			fmt.Println("wrong ip address: " + args[1])
			return

		}
	} else if args[0] == "recive" {

		//launch server
		go server(SETTINGS["port"])

	}
}

func checkParameters(address string) bool {
	if net.ParseIP(address) == nil {
		return false
	}
	return true

}
