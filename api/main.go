package main

import (
	"fmt"

	"github.com/Wraith29/news-app/api/cmd/api"
	"github.com/Wraith29/news-app/api/cmd/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowHeaders = []string{"Authorization", "Content-Type"}

	r.Use(cors.New(corsCfg))

	api.AddApiRoutes(r)

	r.Run(fmt.Sprintf("0.0.0.0:%d", config.Cfg.Port))
}
