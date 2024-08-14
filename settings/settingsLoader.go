package settings

import (
	"encoding/json"
	"io"
	"os"
	"log"
)

const GLOBAL_SETTINGS_FILE = "/etc/lesgofile/lesgofile.json"
const LOCAL_SETTINGS_FILE = "lesgofile.json"

var SETTINGS map[string]string = make(map[string]string)

func setDefaults() {
	SETTINGS["PORT"] = "50000"
	SETTINGS["DIM_BUFFER"] = "1024"
	SETTINGS["DESTINATION_FOLDER"] = "." + string(os.PathSeparator)
	SETTINGS["DISCOVERY_SERVER_PORT"] = ":8828"
	SETTINGS["DISCOVERY_CLIENT_PORT"]= ":8829"
	SETTINGS["DISCOVERY_MAX_TRIES"]="5"
	SETTINGS["BROADCAST_ADDRESS"]="192.168.1.255"
}

// set default settings and try to read config file
func LoadSettings() {
	setDefaults()
	readFile(GLOBAL_SETTINGS_FILE)
	readFile(LOCAL_SETTINGS_FILE)
}
func readFile(name string){

	readFile, err := os.Open(name)
	defer readFile.Close()
	if err == nil {
		byteResult, _ := io.ReadAll(readFile)
		json.Unmarshal([]byte(byteResult), &SETTINGS)
		log.Printf("settings loaded from %s",name)
	}else{
		log.Printf("settings file %s not readed",name)
	}
}
