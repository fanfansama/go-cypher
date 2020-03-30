package models

import (
	"time"
	"miscs"
	"log"
)

type File struct {
	Name 		 string    `json:"name"`
	Size         int64     `json:"size"`
	CreatedAt    time.Time `json:"created_at"`
}

type Files []File


func AllFiles() *Files {
	var files Files= []File{}
//		cars = append(cars, c)
	return &files
}

func NewFile(filename string) *File {
	file := File{Name: filename}
	
	size, err := miscs.FileSize(filename)
	if err != nil {
		log.Print("no file found ("+ filename +")... :" + err.Error())
		file.Size=-1
		return &file
	}
	file.Size = size
	file.CreatedAt = time.Now()
	return &file
}

