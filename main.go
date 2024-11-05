package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AteevMishra/todo-CLI/cmd"
	"github.com/AteevMishra/todo-CLI/db"
)

func main() {
	err := db.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}