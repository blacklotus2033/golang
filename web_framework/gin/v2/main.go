package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		var from = "nobody"
		if p, ok := c.GetQuery("from"); ok {
			from = p
		}
		c.String(200, "double greeting: %s", from)
	})
	go r.Run() //2)不阻塞
	//
	a := gin.New()
	a.GET("/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"in": "admin"})
	})
	a.Run(":8000") //1)用admin来做一个阻塞
}
