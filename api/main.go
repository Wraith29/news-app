package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wraith29/news-app/api/cmd/api"
	"github.com/wraith29/news-app/api/cmd/config"
)

func main() {
	err := config.LoadConfig()

	if err != nil {
		panic(err)
	}
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	api.AddApiRoutes(r)

	r.Run(fmt.Sprintf("0.0.0.0:%d", config.Cfg.Port))
}
