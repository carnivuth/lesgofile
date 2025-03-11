package network

import (
	"encoding/binary"
	"io"
	"net"
	"os"
	"log"
	"path/filepath"
)

// connect to server and upload file
func Sender(address string, port string, name string,buffer_dimension int) {

	//connect to server
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		log.Panicf("could not connect to host %s:%s",address,port)
	}
	log.Printf( "connection succesful to host: %s",address)

	//open file
	info, err := os.Stat(name)
	if err != nil {
		log.Panicf( name,"%s is not a correct path")
	}

	//cheking if file is a directory cause sending folders is not supported
	if info.IsDir() {
		log.Panicf("%s is a directory ",name)
	}

	sendfile(name, conn,buffer_dimension)
	conn.Close()
}
func sendfile(name string, conn net.Conn,buffer_dimension int) {
	defer conn.Close()

	buffer := make([]byte, buffer_dimension)

	//open file to send
	file, err := os.Open(name)
	if err != nil {
		log.Panicf("file %s could not be read",name)
	}
	defer file.Close()

	// get filename
	var _, filename = filepath.Split(name)

	// send name lenght
	binary.Write(conn, binary.BigEndian, int64(len(filename)))

	//send filename
	conn.Write([]byte(filename))

	//send file
	log.Printf( "sending file %s",name)
	for err != io.EOF {
		var n_read int
		n_read, err = file.Read(buffer)
		conn.Write(buffer[:n_read])

	}

}
