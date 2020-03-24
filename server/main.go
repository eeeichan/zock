package main // import "server"

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "ping",
        })
    })
    r.Run(":8888")
  }
