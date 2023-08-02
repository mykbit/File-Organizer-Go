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
	Images    []image
	Documents []document
	Videos    []video
	Audios    []audio
	mutex     sync.Mutex
	wg        sync.WaitGroup
)

type file struct {
	filename    string
	path        string
	size        int64
	lastModDate time.Time
}

type image struct {
	file
}

type document struct {
	file
}

type video struct {
	file
}

type audio struct {
	file
}

func categorizeFile(name string, extension string, category string, filePath string) {
	defer wg.Done()
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
		image := image{file}
		Images = append(Images, image)
	case "Document":
		document := document{file}
		Documents = append(Documents, document)
	case "Video":
		video := video{file}
		Videos = append(Videos, video)
	case "Audio":
		audio := audio{file}
		Audios = append(Audios, audio)
	default:
		panic("Category not found. Please, add it to extensions.json")
	}
	mutex.Unlock()
}

func processFile(name string, extension string, filePath string) {
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
		panic("Extension not found. Please, add it to extensions.json")
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

	// for _, image := range Images {
	// 	pf("%s\n", image.filename)
	// }

	// for _, document := range Documents {
	// 	pf("%s\n", document.filename)
	// }

	// for _, video := range Videos {
	// 	pf("%s\n", video.filename)
	// }

	// for _, audio := range Audios {
	// 	pf("%s\n", audio.filename)
	// }
}
