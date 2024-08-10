package settings

import (
	"encoding/json"
	"io"
	"os"
)

const SETTINGS_FILE = "/etc/lesgofile/lesgofile.json"

var SETTINGS map[string]string = make(map[string]string)

func setDefaults() {
	SETTINGS["PORT"] = "50000"
	SETTINGS["DIM_BUFFER"] = "1024"
	SETTINGS["DESTINATION_FOLDER"] = "." + string(os.PathSeparator)
}

func LoadSettings() {
	// set default settings and try to read config file
	setDefaults()
	readFile, err := os.Open(SETTINGS_FILE)
	defer readFile.Close()
	if err == nil {
		byteResult, _ := io.ReadAll(readFile)
		json.Unmarshal([]byte(byteResult), &SETTINGS)
	}
}
