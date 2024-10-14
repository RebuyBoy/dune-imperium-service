package models

import "io"

type FileData struct {
	Content  io.Reader
	Size     int64
	Filename string
}
