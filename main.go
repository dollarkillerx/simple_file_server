package main

import (
	"github.com/gin-gonic/gin"

	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("./simple_file_server listenAddr listenPath")
	}

	app := gin.Default()

	app.Static("/", os.Args[2])
	log.Println("Listen on: ", os.Args[1])
	if err := app.Run(os.Args[1]); err != nil {
		log.Fatalln(err)
	}
}
