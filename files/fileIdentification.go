package files

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	pf        = fmt.Printf
	Images    []file
	Documents []file
	Videos    []file
	Audios    []file
	mutex     sync.Mutex
	wg        sync.WaitGroup
)

type file struct {
	filename    string
	path        string
	size        int64
	lastModDate time.Time
}

func categorizeFile(name string, extension string, category string, filePath string) {
	pf("Categorizing file %s with extension %s\n", name, extension)
	fileLocation := filePath + "/" + name
	fileInfo, err := os.Stat(fileLocation)
	if err != nil {
		panic(err)
	}
	file := file{name, filePath, fileInfo.Size(), fileInfo.ModTime()}

	mutex.Lock()
	switch category {
	case "Image":
		Images = append(Images, file)
	case "Document":
		Documents = append(Documents, file)
	case "Video":
		Videos = append(Videos, file)
	case "Audio":
		Audios = append(Audios, file)
	default:
		panic("Category not found. Please, add it to extensions.json")
	}
	mutex.Unlock()
}

func processFile(name string, extension string, filePath string) {
	defer wg.Done()
	jsonData, err := os.ReadFile("files/extensions.json")
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
		categorizeFile(name, extension, category, filePath)
	} else {
		pf("Extension %s not found. This file will be skipped. Please, add it to extensions.json\n", extension)
		return
	}
}

// Add file identification methods here
func BrowseFolder(path string) {
	pf("Browsing directory: %s\n", path)
	dir, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	allFiles, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	//exclude hidden files
	var dirEntries []os.FileInfo
	for _, file := range allFiles {
		if file.Name()[0] != '.' {
			dirEntries = append(dirEntries, file)
		}
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			BrowseFolder(path + "/" + entry.Name())
		} else {
			fileName := entry.Name()
			fileExtension := filepath.Ext(fileName)
			wg.Add(1)
			go processFile(fileName, fileExtension, path)
		}
	}

	wg.Wait()

	for _, image := range Images {
		pf("Image: %s\n", image.filename)
	}

	for _, document := range Documents {
		pf("Doc: %s\n", document.filename)
	}

	for _, video := range Videos {
		pf("Video: %s\n", video.filename)
	}

	for _, audio := range Audios {
		pf("Audio: %s\n", audio.filename)
	}
}
