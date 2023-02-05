package network

import (
	"fmt"
	"net"
)

// connect to server and upload file
func Sender(address string, port string) {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connection succesful to host: " + address)
	conn.Close()
}
