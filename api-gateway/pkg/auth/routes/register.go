package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/auth/pb"
)

type RegisterRequestBody struct {
	MSISDN   string `json:"msisdn"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
} //@name RegisterRequestBody

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	b := RegisterRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Msisdn:   b.MSISDN,
		Name:     b.Name,
		Username: b.Username,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
