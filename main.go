package main

import (
	"net/http"
	"time"
	"tokomania/internal/delivery"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	delivery.Routes(r)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}