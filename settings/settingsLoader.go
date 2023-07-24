package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const SETTINGS_FILE = "settings.json"

var SETTINGS map[string]string = make(map[string]string)

func LoadSettings() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	readFile, err := os.Open(exPath + string(os.PathSeparator) + SETTINGS_FILE)

	if err != nil {
		setDefaults()
	} else {
		byteResult, _ := ioutil.ReadAll(readFile)
		json.Unmarshal([]byte(byteResult), &SETTINGS)

	}

	//close file
	readFile.Close()
}

func setDefaults() {

	SETTINGS["PORT"] = "50000"
	SETTINGS["DIM_BUFFER"] = "1024"
	SETTINGS["DESTINATION_FOLDER"] = "./"

}
