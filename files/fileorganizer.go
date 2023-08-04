package fileorganizer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	destpaths "github.com/mykbit/File-Organizer-Go/paths"
)

var (
	pf = fmt.Printf
	wg sync.WaitGroup
)

type file struct {
	fileName  string
	extension string
	path      string
}

func cleanSourceFolder(sourceFolder string) error {
	fmt.Printf("Removing folder %s\n", sourceFolder)
	err := os.RemoveAll(sourceFolder)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("permission denied: unable to clean source folder")
		}
		return err
	}

	// Add a delay and check if the folder still exists
	time.Sleep(time.Millisecond * 5) // Adjust the delay as needed
	if _, err := os.Stat(sourceFolder); err == nil {
		return fmt.Errorf("folder still exists after cleaning")
	}

	return nil
}

func organize(orgFile file, destDirPath string) {
	destFilePath := filepath.Join(destDirPath, orgFile.fileName)
	err := os.Rename(filepath.Join(orgFile.path, orgFile.fileName), destFilePath)
	if err != nil {
		panic(err)
	}
}

func categorizeFile(name string, extension string, category string, filePath string) {
	pf("Categorizing and moving file %s with extension %s\n", name, extension)

	curFile := file{name, extension, filePath}

	switch category {
	case "Image":
		organize(curFile, destpaths.ImagePath)
	case "Document":
		organize(curFile, destpaths.DocumentPath)
	case "Video":
		organize(curFile, destpaths.VideoPath)
	case "Audio":
		organize(curFile, destpaths.AudioPath)
	default:
		panic("Category not found. Please, add it to extensions.json")
	}
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
	go cleanSourceFolder(path)
}
