package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	static_dir := makeStaticDir()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(static_dir))))
	router.Use(mux.CORSMethodMiddleware(router))

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", serverHost, serverPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	InfoLogger.Println("Application Startup Complete")

	InfoLogger.Fatal(srv.ListenAndServe())

}
