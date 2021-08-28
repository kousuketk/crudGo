# crudGo

### start

- docker-compose
```
$ docker compose up -d
```

- run server
```
$ docker exec -it crudgo_go_1 bash
# go run main.go
```

- look mysql
```
docker exec -it crudgo_db_1 bash
```

### endpoints
```
$ curl http://localhost:8080/api/v1/users | jq .
$ curl http://localhost:8080/api/v1/users/{user_id} | jq .
$ curl -X POST -H 'Content-Type:application/json' -d '{"name":"test_name", "self_introduction":"test_self_introduction", "email":"testuser@test.com", "password_digest":"test123", "address":"test_address", "phone_number":"00000000000"}' http://localhost:8080/api/v1/me/ | jq .

// login(get jwt)
$ curl http://localhost:8080/api/login -X POST -H 'Content-Type:application/json' -d '{"email":"testuser@test.com", "password":"test123"}'
// with jwt
$ curl http://localhost:8080/api/v1/me/
$ curl -X POST -H 'Content-Type:application/json' -d '{"name":"test_name", "self_introduction":"test_self_introduction", "email":"testuser@test.com", "password_digest":"test123", "address":"test_address", "phone_number":"00000000000"}' http://localhost:8080/api/v1/me/ | jq .
$ curl http://localhost:8080/api/v1/me/ -X DELETE
$ curl http://localhost:8080/api/v1/logout
```

### test
```
go test ./test/controller/ -cover
```
