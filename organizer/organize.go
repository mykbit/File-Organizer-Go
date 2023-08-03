package organizer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"runtime"
)

var (
	pf           = fmt.Printf
	imagePath    string
	documentPath string
	audioPath    string
	videoPath    string
)

func SetDefaultDestinationPaths() {
	file, err := os.Open("organizer/baseDirConfig.json")
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

	imagePath = os["base"].(string) + user.Username + os["folders"].(map[string]interface{})["Images"].(string)
	documentPath = os["base"].(string) + user.Username + os["folders"].(map[string]interface{})["Documents"].(string)
	audioPath = os["base"].(string) + user.Username + os["folders"].(map[string]interface{})["Audios"].(string)
	videoPath = os["base"].(string) + user.Username + os["folders"].(map[string]interface{})["Videos"].(string)

	pf("Image path: %s\n", imagePath)
	pf("Document path: %s\n", documentPath)
	pf("Audio path: %s\n", audioPath)
	pf("Video path: %s\n", videoPath)
}
