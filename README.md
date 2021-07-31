- build, start
docker compose build
docker compose up -d

- bashに入る, 起動
docker compose exec go bash
go run main.go

- mysql
docker exec -it sampledocker_db_1 bash