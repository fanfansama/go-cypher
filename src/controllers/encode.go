package controllers

import (
	"log"
	"runtime"
	"sync"

	"net/http"

	"cypher"
	"miscs"
)

func Encode(w http.ResponseWriter, r *http.Request) {

	handler, err := miscs.Upload(w, r)

	if err != nil {
		log.Fatal("cannot upload file, encryption process failed")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Maximum process to run based on cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Waitgroup for multiple process
	var wg sync.WaitGroup

	// FIXME: delete dest file before to encode
	err0 := cypher.EncryptFile("./in/"+handler, "./out-enc/"+handler+".enc", key)
	
	if err0 != nil {
		log.Print("cannot decrypt file")
		log.Print(err0)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return	
	} else {
		log.Print("encrypted To: " + handler+".enc")
	}

	// Wait till every goroutine has finished
	wg.Wait()

	miscs.DeleteFile("./in/"+handler)
	FilesShow(w, r, "files/encoded/"+handler+".enc") 

}