package main

import (
	"github.com/franszhafran/sahamify-be/modules/authentication/presentation"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authentication := presentation.Routes{}
	authentication.Register(r)

	// Run the server
	r.Run(":9200")
}
