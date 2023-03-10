package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/OlegDjur/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type LoginResponseBody struct {
// 	Status string `json:"status"`
// 	token  string `json:"token"`
// }

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	fmt.Println(res.Token)

	ctx.JSON(http.StatusOK, &res)
}
