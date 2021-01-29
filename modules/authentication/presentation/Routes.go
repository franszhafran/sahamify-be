package presentation

import (
	"github.com/gin-gonic/gin"
)

// Routes is a facade for registering routes in this modules
type Routes struct{}

// Register the routes
func (r Routes) Register(e *gin.Engine) {
	e.GET("/login", Login)
}
