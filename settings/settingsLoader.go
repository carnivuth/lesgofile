package settings

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const SETTINGS_FILE = "settings.conf"

var SETTINGS map[string]string = make(map[string]string)

func LoadSettings() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	readFile, err := os.Open(exPath + SETTINGS_FILE)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	//load settings
	for fileScanner.Scan() {
		setting := strings.Split(fileScanner.Text(), "=")
		SETTINGS[setting[0]] = setting[1]
	}
	//close file
	readFile.Close()
}
