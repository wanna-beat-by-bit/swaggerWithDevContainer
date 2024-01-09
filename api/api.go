package api

import (
	"io"
	"log"
	"net/http"
	"os"
)

// TestHandler example
//
//	@Summary		returns file
//	@Description	getting file
//	@ID				crop
//	@Produce		png
//	@Router			/testapi/ [get]
//	@Success		200		{file}	png	"![one file](test.png)"
//	@Success		200		{string}	string	"hello world"
//	@Header 	    200 {string} SuperMeta "i am super meta data"
func TestHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("test.png")
	if err != nil {
		log.Fatalf("opening file: %s", err.Error())
	}

	textFile, err := io.ReadAll(file)
	w.Header().Set("SuperMeta", "i am super meta data")
	if _, err := w.Write(textFile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
