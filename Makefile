postgres:
	docker run --name=postgres -e POSTGRES_USER=olegdjur -e POSTGRES_PASSWORD=qwerty -p 5432:5432 -d postgres

createdb:
	docker exec -it postgres createdb --username=olegdjur --owner=olegdjur auth_svc \
	docker exec -it postgres createdb --username=olegdjur --owner=olegdjur order_svc \
	docker exec -it postgres createdb --username=olegdjur --owner=olegdjur product_svc