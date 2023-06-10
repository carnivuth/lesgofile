package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"github.com/carnivuth/lesgofile/settings"
)

// connect to server and upload file
func Sender(address string, port string, name string) {

	dim, err := strconv.Atoi(settings.SETTINGS["DIM_BUFFER"])
	buffer := make([]byte, dim)
	//connect to server
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connection succesful to host: " + address)
	//open file
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
	conn.Close()
}
