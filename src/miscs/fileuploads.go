package miscs

import (
	"os"
	"io"
	"log"
	"net/http"

)

func Upload(w http.ResponseWriter, r *http.Request) (response string, err error) {

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer file.Close()
	
	DeleteFile("./in/"+handler.Filename)

	f, err := os.OpenFile("./in/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)
	log.Println("Received: " + handler.Filename)
	return handler.Filename, nil
}