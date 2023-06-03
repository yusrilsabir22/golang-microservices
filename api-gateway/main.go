package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yusrilsabir22/orderfaz/api-gateway/docs"

	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/auth"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/config"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/logistic"
)

// @title API Service
// @version 1.0
// @description API in go using Gin framework
// @host localhost:3000
// @BasePath /auth
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authSvc := *auth.RegisterRoutes(r, &c)
	logistic.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
