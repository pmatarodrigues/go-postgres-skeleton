package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/pmatarodrigues/go-starter/server"

	"github.com/pmatarodrigues/go-starter/storage"

	"github.com/pmatarodrigues/go-starter/config"
	"github.com/pmatarodrigues/go-starter/connections"

	"github.com/gin-gonic/gin"
)

func main() {
	// Console color not needed when writing to file
	gin.DisableConsoleColor()

	// Log to file
	f, _ := os.Create("go-starter.log")
	gin.DefaultWriter = io.MultiWriter(f)

	if err := start(); err != nil {
		panic(fmt.Sprintf("Cannot start server. %s", err))
	}

	// r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run()
}

func start() error {
	// Load environment data and config file
	config := config.New()

	db, err := connections.NewDatabase(config.Database)
	if err != nil {
		return err
	}

	// Save DB data to repository
	repository := storage.NewRepo(db)

	// Concurrent handle graceful server shutdown
	go handleShutdown(db)

	s := server.New(repository)
	if err = s.Run(fmt.Sprintf(":%d", config.Server.Port)); err != nil {
		return err
	}

	fmt.Print("Server running on port %d \n", config.Server.Port)

	return nil
}

func handleShutdown(db *sql.DB) {
	// Listen for CTRL+C input
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	// Handle close server

	if err := db.Close(); err != nil {
		fmt.Printf("Error trying to close connection to the database. %v", err)
	}
	os.Exit(0)
}
