package api

import "io"

type PlayerCreateRequest struct {
	Nickname string `json:"nickname"`
	Avatar   Avatar `json:"-"`
}

type Avatar struct {
	File     io.Reader
	Size     int64
	Filename string
}
