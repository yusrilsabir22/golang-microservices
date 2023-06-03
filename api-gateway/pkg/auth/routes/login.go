package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
	MSISDN   string `json:"msisdn"`
	Password string `json:"password"`
} //@name LoginRequestBody

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Msisdn:   b.MSISDN,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	fmt.Println(res.Message)
	ctx.JSON(int(res.Status), &res)
}
