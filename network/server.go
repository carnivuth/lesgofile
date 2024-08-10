package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/carnivuth/lesgofile/logger"
	"github.com/carnivuth/lesgofile/settings"
)

//wait for client and download file

func Reciver(port string) {

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic("unable to listen over port: " + port)
	}
	//loop
	logger.Emit(logger.Log, "start listen on port: "+port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Emit(logger.Log, "unable to accept request from client")
		} else {
			//handle connection with go routine
			go saveFile(conn)

		}
	}

}
func saveFile(conn net.Conn) {
	var size int64
	var n_read int
	dim, err := strconv.Atoi(settings.SETTINGS["DIM_BUFFER"])
	buffer := make([]byte, dim)
	filename := make([]byte, dim)
	logger.Emit(logger.Log, "connected to client")
	if err != nil {
		fmt.Println(err)
		return
	}

	binary.Read(conn, binary.BigEndian, &size)

	//read filename
	n_read, err = conn.Read(filename[:size])
	logger.Emit(logger.Log, "recived "+string(filename[:n_read])+" filename\n")

	//create file
	file, err := os.Create(settings.SETTINGS["DESTINATION_FOLDER"] + string(filename[:n_read]))
	if err != nil {
		panic(err)
	}
	//write file
	for err != io.EOF {
		n_read, err = conn.Read(buffer)
		file.Write(buffer[:n_read])
	}
	file.Close()
	logger.Emit(logger.Log, "connection terminated ")

}
