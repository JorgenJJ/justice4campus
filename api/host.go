package api

import (
	"github.com/gin-gonic/gin"
)

func TestPost(c *gin.Context) {
	c.JSON(200, gin.H{"status": "success"})

}
