package destpaths

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"runtime"
)

var (
	pf           = fmt.Printf
	ImagePath    string
	DocumentPath string
	AudioPath    string
	VideoPath    string
)

func SetDefaultDestinationPaths() {
	file, err := os.Open("paths/baseDirConfig.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config map[string]map[string]interface{}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	os := config[runtime.GOOS]

	ImagePath = os["base"].(string) + user.Username + os["folders"].(map[string]interface{})["Images"].(string)
	DocumentPath = os["base"].(string) + user.Username + os["folders"].(map[string]interface{})["Documents"].(string)
	AudioPath = os["base"].(string) + user.Username + os["folders"].(map[string]interface{})["Audios"].(string)
	VideoPath = os["base"].(string) + user.Username + os["folders"].(map[string]interface{})["Videos"].(string)
}
