package middleware
import (
	"net/http" 
	"github.com/gin-gonic/gin"
	"log"
	_ "github.com/lib/pq"
)

func Auth(c *gin.Context) {
	log.Println("start middleware")
	authKey := c.GetHeader("Authorization")
	if authKey != "Bearer token123" {
		c.JSON(http.StatusUnauthorized,gin.H{"error": "don't have the permission!"})
		c.Abort()
		return 
	}
	c.Next()
	log.Println("end middleware")

}