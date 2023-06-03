package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/yusrilsabir22/orderfaz/api-gateway/docs"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/auth/routes"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/config"
	_ "github.com/yusrilsabir22/orderfaz/logistic-svc/docs"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	routes.POST("/validate", svc.Validate)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Validate(ctx *gin.Context) {
	routes.Validate(ctx, svc.Client)
}
