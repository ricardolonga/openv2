package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Olá, Longa!")
	})

	bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	err := http.ListenAndServe(bind, r)
	if err != nil {
		panic(err)
	}
}
