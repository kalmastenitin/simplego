package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	dbUsername = "root"
	dbPassword = "password"
	dbHost = "127.0.0.1"
	dbPort = "3306"
	serverHost = "127.0.0.1"
	serverPort = "8001"
)

func getDSNinfo() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUsername, dbPassword, dbHost, dbPort)
}


func main(){
	var dir string

	dbname := "gotestdb"

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	db, err := sql.Open("mysql", getDSNinfo())

	if err != nil {
        log.Printf("Error %s when connecting DB\n", err)
        return
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()

    _, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
        log.Printf("Error %s when creating DB\n", err)
        return
	}
    
	fmt.Println("connected to database ...")
    
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", serverHost, serverPort),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server startup completed")

	log.Fatal(srv.ListenAndServe())
	defer db.Close()
	
}