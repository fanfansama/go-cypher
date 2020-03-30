package controllers

import (
	"log"
	"runtime"
	"sync"

	"net/http"

	"cypher"
	"miscs"
)

func Decode(w http.ResponseWriter, r *http.Request) {

	handler, err := miscs.Upload(w, r)

	if err != nil {
		log.Print("cannot upload file, encryption process failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Maximum process to run based on cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Waitgroup for multiple process
	var wg sync.WaitGroup

	// FIXME: delete dest file before to encode
	err0 := cypher.DecryptFile("./in/"+handler, "./out-dec/"+handler+".dec", key)
	
	if err0 != nil {
		log.Print("cannot decrypt file")
		log.Print(err0)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else {
		log.Print("decrypter To: " + handler+".dec")
	}

	// Wait till every goroutine has finished
	wg.Wait()

	miscs.DeleteFile("./in/"+handler)
	FilesShow(w, r, "files/decoded/"+handler+".dec") 

}