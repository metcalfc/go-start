package main

import "rsc.io/quote"
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, quote.Hello())
	})

	r.Run(":3000")
}