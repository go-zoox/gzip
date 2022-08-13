package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-zoox/gzip"
	"github.com/go-zoox/zoox"
)

func main() {
	r := zoox.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Get("/ping", func(c *zoox.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
