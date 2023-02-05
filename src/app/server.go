package lesGoFile

import (
	"fmt"
	"net"
)

//wait for client and download file

func server(port string) {

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("unable to listen over port: " + port)
	}
	//loop
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("unable to accept request from client")
		} else {
			//handle connection with go routine
			go saveFile(&conn)
		}
	}

}
func saveFile(conn *net.Conn) {
	fmt.Println("connected to client")

	///create protocol for transfer file over server

}
