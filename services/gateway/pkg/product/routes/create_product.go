package routes

import (
	"context"
	"net/http"

	"github.com/OlegDjur/go-grpc-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	Name  string `jaon:"name"`
	Stock int64  `json:"stok"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
