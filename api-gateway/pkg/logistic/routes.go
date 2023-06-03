package logistic

import (
	"github.com/gin-gonic/gin"
	_ "github.com/yusrilsabir22/orderfaz/api-gateway/docs"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/auth"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/config"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/logistic/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/logistic")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateLogistic)
	routes.GET("/", svc.FindLogistic)
}

func (svc *ServiceClient) CreateLogistic(ctx *gin.Context) {
	routes.CreateLogistic(ctx, svc.Client)
}

func (svc *ServiceClient) FindLogistic(ctx *gin.Context) {
	routes.FindLogistic(ctx, svc.Client)
}
