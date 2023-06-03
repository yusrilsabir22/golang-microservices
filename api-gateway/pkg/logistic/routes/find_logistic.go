package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/logistic/pb"
)

type FindLogisticQuery struct {
	DestinationName string `json:"destination_name"`
	OriginName      string `json:"origin_name"`
}

func FindLogistic(ctx *gin.Context, c pb.LogisticServiceClient) {
	b := FindLogisticQuery{}

	if err := ctx.BindQuery(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneLogisticRequest{
		DestinationName: b.DestinationName,
		OriginName:      b.OriginName,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
