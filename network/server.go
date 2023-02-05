package network

import (
	"fmt"
	"net"
)

//wait for client and download file

func Reciver(port string, terminate chan int) {

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("unable to listen over port: " + port)
	}
	//loop
	fmt.Println("start listen on port: " + port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("unable to accept request from client")
		} else {
			//handle connection with go routine
			go saveFile(&conn, terminate)
			<-terminate
			fmt.Println("connection terminated ")
		}
	}

}
func saveFile(conn *net.Conn, terminate chan int) {
	fmt.Println("connected to client")
	terminate <- 1
	///create protocol for transfer file over server

}
