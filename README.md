- build, start
docker compose build
docker compose up -d

- bashに入る, 起動
docker exec -it crudgo_go_1 bash
go run main.go

- mysql
docker exec -it crudgo_db_1 bash