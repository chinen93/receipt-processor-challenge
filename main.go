package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	outfile, _ = os.Create("project.log")
	l          = log.New(outfile, "", 0)
)

func main() {
	router := gin.Default()
	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getReceiptPoints)

	l.SetFlags(log.LstdFlags | log.Lmicroseconds)
	l.Println("Server Starts")

	router.Run("localhost:8080")
}
