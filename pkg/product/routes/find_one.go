package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/OlegDjur/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.FindeOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
