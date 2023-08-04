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

func getDestinationFolders(r *bufio.Reader, s string) string {
	pf("%s", s)
	input, err := r.ReadString('\n')
	if err != nil {
		pf("Error: %v\nTry again: ", err)
		getDestinationFolders(r, s)
	}
	input = strings.TrimSpace(input)
	_, err = os.Stat(input)
	if err != nil {
		pf("Error: %v\nTry again ", err)
		getDestinationFolders(r, s)
	}
	return input
}

func getChoice(r *bufio.Reader) {
	choice, err := r.ReadString('\n')
	if err != nil {
		pf("Error: %v\nTry again: ", err)
		getChoice(r)
	}
	choice = strings.TrimSpace(choice)
	if choice != "y" && choice != "n" {
		pf("Error: Invalid choice\nTry again: ")
		getChoice(r)
	}
	if choice == "y" {
		pf("\nPlease, specify the destination folder for each category\n")
		destpaths.ImagePath = getDestinationFolders(r, "Images: ")
		destpaths.DocumentPath = getDestinationFolders(r, "Documents: ")
		destpaths.AudioPath = getDestinationFolders(r, "Audios: ")
		destpaths.VideoPath = getDestinationFolders(r, "Videos: ")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pf("Welcome to File Organizer!\nPlease, choose the directory you want to organize: ")
	sourceFolder := getSourceFolder(reader)
	destpaths.SetDefaultDestinationPaths()
	pf("\nDefault destination folders:\n *Images: %s\n *Documents: %s\n *Audios: %s\n *Vidoes: %s\n",
		destpaths.ImagePath, destpaths.DocumentPath, destpaths.AudioPath, destpaths.VideoPath)
	pf("\nWould you like to specify the destination folders? (y/n): ")
	getChoice(reader)
	fileorganizer.BrowseFolder(sourceFolder)
}
