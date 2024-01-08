package api

import "net/http"

// TestHandler example
//
//		@Summary		outputs test message
//		@Description	get string
//		@ID				crop
//		@Produce		plain
//		@Success		200		{file}	octet-stream			"./myfile.txt"
//		@Router			/testapi [get]
//	 	@Header 	    200 {string} SuperMeta "i am super meta data"
func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("SuperMeta", "i am super meta data")
	if _, err := w.Write([]byte("Hello Postman!")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
