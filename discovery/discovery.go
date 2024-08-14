package discovery

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"time"
)

const DISCOVERY_SERVER_PORT= ":8828"
const DISCOVERY_CLIENT_PORT= ":8829"
const DISCOVERY_MAX_TRIES=5

// structs for request and response
type DiscoveryRequest struct{
	Address string
}
type DiscoveryResponse struct{
	Address string
	Hostname string
}

func Send_discovery_request() [] DiscoveryResponse{

	var available_servers [] DiscoveryResponse
	pc, err := net.ListenPacket("udp4",DISCOVERY_CLIENT_PORT)
	if err != nil {
		log.Panic("unable to create udp broadcast packet")
	}
	defer pc.Close()

	addr,err := net.ResolveUDPAddr("udp4", "192.168.1.255"+ DISCOVERY_SERVER_PORT)
	if err != nil {
		log.Panicf("unable to create broadcast address, err %s",err)
	}

	//prepare broadcast address
	var request DiscoveryRequest
	request.Address= GetOutboundIP().String()
	json_request, err := json.Marshal(request)
	if err != nil {
		log.Panic("unable to marshal udp discovery request")
	}

	_,err = pc.WriteTo([]byte(json_request), addr)
	if err != nil {
		log.Panic("unable to send udp discovery request")
	}
	// read response from server
	buf := make([]byte, 1024)
	timeout:= time.Now().Local().Add(time.Second * time.Duration(5))
	pc.SetReadDeadline(timeout)
	for i := 0; i< DISCOVERY_MAX_TRIES; i++ {

		n,_,err := pc.ReadFrom(buf)
		if err == nil {
			if (n> 0){
				var response DiscoveryResponse
				err = json.Unmarshal(buf[:n],&response)
				if err != nil {
					log.Panic("unable to unmarshal discovery response")
				}
				available_servers=append(available_servers,response)
			}
		}
	}
	return available_servers
}

func Listen_discovery_requests() {
	pc,err := net.ListenPacket("udp4", DISCOVERY_SERVER_PORT)
	if err != nil {
		log.Panic("unable to listen for discovery requests")
	}
	log.Printf("listening for discovery requests on port %s", DISCOVERY_SERVER_PORT)

	for {
		buf := make([]byte, 1024)
		_,addr,err := pc.ReadFrom(buf)
		if err != nil {
			log.Panic("Unable to read packet")
		}

		go send_discovery_response(pc, addr)
	}
}

func send_discovery_response(pc net.PacketConn,addr net.Addr) {

	var response DiscoveryResponse
	response.Address = GetOutboundIP().String()
	name,err := os.Hostname()
	response.Hostname = name
	if err != nil {
		log.Panic("unable to get hostname")
	}

	json_response, err := json.Marshal(response)
	if err != nil {
		log.Panic("unable to marshal udp discovery response")
	}

	log.Printf("sending response to %s", addr)
	_,err = pc.WriteTo([]byte(json_response), addr)
	if err != nil {
		log.Panicf("unable to send udp discovery response err: %s",err)
	}
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Panic("unable to get the preferred local ip address")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
