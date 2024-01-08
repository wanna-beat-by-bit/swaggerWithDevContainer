package main

import (
	"net/http"

	"github.com/rs/cors"

	"swaggerTest/api"
)

//	@title			Swagger Example API
//	@version		2.0
//	@description	This is a sample server for company.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support //	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// 	@schemes	http
//	@host		localhost:8573
//	@BasePath	/

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/testapi", api.TestHandler)
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs"))))
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}).Handler(mux)
	// serve swagger docs
	http.ListenAndServe(":8573", handler)
}
