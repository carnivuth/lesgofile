package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/carnivuth/lesgofile/settings"
)

//wait for client and download file

func Reciver(port string, terminate chan int) {

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("unable to listen over port: " + port)
	}
	//loop
	fmt.Println("start listen on port: " + port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("unable to accept request from client")
		} else {
			//handle connection with go routine
			go saveFile(conn, terminate)
			<-terminate

		}
	}

}
func saveFile(conn net.Conn, terminate chan int) {
	var size int64
	var n_read int
	dim, err := strconv.Atoi(settings.SETTINGS["DIM_BUFFER"])
	buffer := make([]byte, dim)
	filename := make([]byte, dim)
	fmt.Println("connected to client")
	if err != nil {
		fmt.Println(err)
		return
	}

	binary.Read(conn, binary.BigEndian, &size)

	//read filename
	n_read, err = conn.Read(filename[:size])
	fmt.Printf("recived %s filename\n", string(filename[:n_read]))

	//create file
	file, err := os.Create(settings.SETTINGS["DESTINATION_FOLDER"] + string(filename[:n_read]))
	if err != nil {
		fmt.Println(err)
		return
	}
	//write file
	for err != io.EOF {
		n_read, err = conn.Read(buffer)
		file.Write(buffer[:n_read])
	}
	file.Close()
	fmt.Println("connection terminated ")
	terminate <- 1

}
