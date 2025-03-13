package settings

import (
  "encoding/json"
  "io"
  "os"
  "log"
  "errors"
)

const LOCAL_SETTINGS_FILE = "lesgofile.json"

type settings struct {
  File_server_port string
  Buffer_dimension int
  Destination_folder string
  Discovery_server_port string
  Discovery_client_port string
  Discovery_max_tries int
  Broadcast_address string
}

func defaultSettings() settings {

  settings := settings{
    File_server_port : "50000",
    Buffer_dimension : 1024,
    Destination_folder: "." + string(os.PathSeparator),
    Discovery_server_port: "8828",
    Discovery_client_port : "8829",
    Discovery_max_tries: 5,
    Broadcast_address: "255.255.255.255",
  }

  return settings
}

func readSettingsFile(name string) (settings,error){

  var settings settings
  readFile, err := os.Open(name)
  defer readFile.Close()
  if err == nil {
    byteResult, _ := io.ReadAll(readFile)
    json.Unmarshal([]byte(byteResult), &settings)
    log.Printf("settings loaded from %s",name)
    return settings, nil
  }else{
    log.Printf("settings file %s not readed",name)
    return settings,errors.New("settings file not readed")
  }
}

// set default settings and try to read config file
func LoadSettings() settings{
  var settings = defaultSettings()
  var settings_from_file,err= readSettingsFile(LOCAL_SETTINGS_FILE)
  if err == nil{
    settings = settings_from_file
  }
  return settings
}
func DumpSettings()([]byte,error){
  return json.Marshal(defaultSettings())
}
