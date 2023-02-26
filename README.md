# Microservices in Go with gRPC, \n API Gateway, and Authentication

Это простой пример интернет магазина

## Technologies used

- [gRPC](https://github.com/grpc/grpc-go).
- [Gin Framework](https://github.com/gin-gonic/gin).
- [postgreSQL](https://www.postgresql.org).
- [Docker](https://www.docker.com/).
- [viper](https://github.com/spf13/viper).
- [JWT](github.com/golang-jwt/jwt).

Проект состоит из 3 микросервисов и 1 API Gateway, который будет обрабатывать
входящие HTTP-запросы. HTTP-запросы перенаправляются на микросервисы с помощью gRPC.


`         `<img align="center" width="686" height="601" src="https://github.com/OlegDjur/Readme/blob/master/grpc_shop/simple_shop_drawio.png">