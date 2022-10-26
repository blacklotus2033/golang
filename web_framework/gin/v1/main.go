package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		var from = "nobody"
		if p, ok := c.GetQuery("from"); ok {
			from = p
		}
		c.String(200, "greeting: %s", from)
	})
	r.Run(":8080") //1)默认应该是8080
}
