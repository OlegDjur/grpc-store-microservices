package main

import (
	"log"

	"github.com/OlegDjur/go-grpc-api-gateway/pkg/auth"
	"github.com/OlegDjur/go-grpc-api-gateway/pkg/config"
	"github.com/OlegDjur/go-grpc-api-gateway/pkg/order"
	"github.com/OlegDjur/go-grpc-api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
