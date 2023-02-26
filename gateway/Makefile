postgres:
	docker run --name=postgres -e POSTGRES_USER=olegdjur -e POSTGRES_PASSWORD=qwerty -p 5432:5432 -d postgres

createdb:
	docker exec -it postgres createdb --username=olegdjur --owner=olegdjur auth_svc \
	docker exec -it postgres createdb --username=olegdjur --owner=olegdjur order_svc \
	docker exec -it postgres createdb --username=olegdjur --owner=olegdjur product_svc

proto:
	protoc pkg/**/pb/*.proto --go_out=.
	protoc pkg/**/pb/*.proto --go-grpc_out=.

server:
	go run cmd/main.go