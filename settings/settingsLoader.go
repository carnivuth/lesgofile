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
		panic(err)
	}
	byteResult, _ := ioutil.ReadAll(readFile)
	json.Unmarshal([]byte(byteResult), &SETTINGS)

	/*fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	//load settings
	for fileScanner.Scan() {
		setting := strings.Split(fileScanner.Text(), "=")
		SETTINGS[setting[0]] = setting[1]
	}*/
	//close file
	readFile.Close()
}
