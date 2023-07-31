package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var pf = fmt.Printf
var Images []image
var Documents []document
var Videos []video
var Audios []audio

type file struct {
	filename string
	path     string
	size     int64
}

type image struct {
	file
	createdDate time.Time
	width       int
	height      int
}

type document struct {
	file
}

type video struct {
	file
	duration int
	width    int
	height   int
}

type audio struct {
	file
	duration int
}

func processFile(name string, extension string) {
	jsonData, err := ioutil.ReadFile("extensions.json")
	if err != nil {
		panic(err)
	}

	var categoryMap map[string]string
	err = json.Unmarshal(jsonData, &categoryMap)
	if err != nil {
		panic(err)
	}

	category, found := categoryMap[extension]
	if found {
		pf("Found file %s with extension %s\n", name, extension)
		//categorizeFile(name, extension, category)
	} else {
		panic("Extension not found. Please, add it to extensions.json")
	}
}

// Add file identification methods here
func BrowseFolder(path string) {
	pf("Browsing directory: %s\n", path)
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, entry := range dirEntries {
		if entry.IsDir() {
			BrowseFolder(path + "/" + entry.Name())
		} else {
			fileName := entry.Name()
			fileExtension := filepath.Ext(fileName)
			processFile(fileName, fileExtension)
		}
	}
}
