package settings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const SETTINGS_FILE = "settings.conf"

var SETTINGS map[string]string = make(map[string]string)

func LoadSettings() {

	readFile, err := os.Open(SETTINGS_FILE)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	//load settings
	for fileScanner.Scan() {
		setting := strings.Split(fileScanner.Text(), ":")
		SETTINGS[setting[0]] = setting[1]
	}
	//close file
	readFile.Close()
}
