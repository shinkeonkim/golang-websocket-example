package network

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Network struct {
	engin *gin.Engine
}

func NewServer() *Network {
	n := &Network{engin: gin.New()}

	n.engin.Use(gin.Logger())   // Log all requests to the console
	n.engin.Use(gin.Recovery()) // Recover from any panics
	n.engin.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	})) // CORS Settings

	return n
}

func (n *Network) StartServer() error {
	log.Println("Starting Server")

	return n.engin.Run(":8080")
}
