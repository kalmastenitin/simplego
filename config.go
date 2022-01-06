package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	dbUsername = "root"
	dbPassword = "toolkitsecret"
	dbHost     = "127.0.0.1"
	dbPort     = "3306"
	serverHost = "127.0.0.1"
	serverPort = "8001"
	dbname     = "gotestdb"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func getDSNinfo() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUsername, dbPassword, dbHost, dbPort)
}

func dbConnection() {
	db, err := sql.Open("mysql", getDSNinfo())

	if err != nil {
		InfoLogger.Printf("Error %s when connecting DB\n", err)
		return
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		InfoLogger.Printf("Error %s when creating DB\n", err)
		return
	}

	InfoLogger.Println("connected to database...")
	defer db.Close()
}

func makeStaticDir() string {
	current_directory, _ := os.Getwd()
	static_dir := filepath.Join(current_directory, "static_dir")
	if _, err := os.Stat(static_dir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(static_dir, 0755)
		}
	}
	return static_dir
}

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	dbConnection()
}
