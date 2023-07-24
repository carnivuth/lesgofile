package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"github.com/carnivuth/lesgofile/logger"
	"github.com/carnivuth/lesgofile/settings"
)

// connect to server and upload file
func Sender(address string, port string, name string) {

	//connect to server
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.Emit(logger.Log, "connection succesful to host: "+address)
	//open file

	info, err := os.Stat(name)
	if err != nil {
		logger.Emit(logger.Log, name+" is not a correct path ")
		return
	}

	if info.IsDir() {
		logger.Emit(logger.Log, name+" is a directory ")
		//TODO directory support
	} else {
		logger.Emit(logger.Log, "sending "+name)
		sendfile(name, conn)
	}

	conn.Close()
}
func sendfile(name string, conn net.Conn) {
	dim, err := strconv.Atoi(settings.SETTINGS["DIM_BUFFER"])
	buffer := make([]byte, dim)
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get filename
	var _, filename = filepath.Split(name)

	// send name lenght
	binary.Write(conn, binary.BigEndian, int64(len(filename)))

	//send filename
	conn.Write([]byte(filename))

	//send file
	for err != io.EOF {
		var n_read int
		n_read, err = file.Read(buffer)
		conn.Write(buffer[:n_read])

	}
	file.Close()

}
