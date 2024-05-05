package handlers

import "github.com/gin-gonic/gin"

func (app *MyApp) Routes(r *gin.Engine) {
	r.GET("/", Health)
}

// health function to check the health (Server is running or not)
func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"error":   false,
		"message": "check health of the app",
	})
}
