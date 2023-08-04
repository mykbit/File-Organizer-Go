package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	fileorganizer "github.com/mykbit/File-Organizer-Go/files"
	destpaths "github.com/mykbit/File-Organizer-Go/paths"
)

var pf = fmt.Printf

func getSourceFolder(r *bufio.Reader) string {
	sourceFolder, err := r.ReadString('\n')
	if err != nil {
		pf("Error: %v\nTry again: ", err)
		return getSourceFolder(r)
	}
	sourceFolder = strings.TrimSpace(sourceFolder)
	_, err = os.Stat(sourceFolder)
	if err != nil {
		pf("Error: %v\nTry again: ", err)
		return getSourceFolder(r)
	}
	return sourceFolder
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pf("Welcome to File Organizer!\nPlease, choose the directory you want to organize: ")
	sourceFolder := getSourceFolder(reader)
	destpaths.SetDefaultDestinationPaths()
	fileorganizer.BrowseFolder(sourceFolder)
}
