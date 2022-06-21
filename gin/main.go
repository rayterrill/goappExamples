package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//use the gin recovery middleware
	router.Use(gin.Recovery())

	//build some routes
	router.GET("/ping", ping)
	router.GET("/slow", doSomethingSlow)

	//build a server so we can control the settings
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s", err)
		}
	}()

	//wait for interrupt signal - 5 sec timeout
	quit := make(chan os.Signal, 1)
	//remember - we cant catch SIGTERM - no need to add that
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//wait from signal
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	//catch ctx.Done() - timeout exceeded
	<-ctx.Done()
	log.Println("Timeout exceeded - Server exiting")
}

func ping(c *gin.Context) {
	fmt.Println("server: ping handler started")
	defer fmt.Println("server: ping handler ended")

	c.JSON(http.StatusOK, gin.H{
		"message": "PONG",
	})
}

func doSomethingSlow(c *gin.Context) {
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(c.Writer, "hello\n")
	case <-c.Done():
		err := c.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(c.Writer, err.Error(), internalError)
	}
}
