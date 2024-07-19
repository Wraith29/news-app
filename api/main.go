package main

import (
	"fmt"
	"os"
	"rss-app/cmd/api"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultPort = 2912

func getPort() int {
	envPort := os.Getenv("RSS_APP_PORT")

	if len(envPort) == 0 {
		return defaultPort
	}

	port, err := strconv.Atoi(envPort)

	if err != nil {
		return defaultPort
	}

	return port
}

func main() {
	port := getPort()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	api.AddApiRoutes(r)

	r.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
