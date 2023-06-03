package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/logistic/pb"
)

type CreateLogisticRequestBody struct {
	LogisticName    string `json:"logistic_name"`
	Amount          int    `json:"amount"`
	DestinationName string `json:"destination_name"`
	OriginName      string `json:"origin_name"`
	Duration        string `json:"duration"`
}

func CreateLogistic(ctx *gin.Context, c pb.LogisticServiceClient) {
	b := CreateLogisticRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// userId, _ := ctx.Get("userId")

	res, err := c.CreateLogistic(context.Background(), &pb.CreateLogisticRequest{
		LogisticName:    b.LogisticName,
		Amount:          int64(b.Amount),
		DestinationName: b.DestinationName,
		OriginName:      b.OriginName,
		Duration:        b.Duration,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
