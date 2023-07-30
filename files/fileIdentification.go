package files

import "time"

type File struct {
	filename string
	path     string
	size     int64
}

type Image struct {
	File
	createdDate time.Time
	width       int
	height      int
}

type Document struct {
	File
}

type Video struct {
	File
	duration int
	width    int
	height   int
}

type Music struct {
	File
	duration int
}

// Add file identification methods here
