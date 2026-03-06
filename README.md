# go-auth


## Sync dependencies
     
```bash
go mod tidy
```     

## Tools

Create a hash password

```bash
go run hash_password.go "123456Abcde"
```

```bash
go build -o ./dist/hash_password hash_password.go
```

```bash
./dist/hash_password "123456Abcde"
```

## Docker build

```bash
docker build -t go-auth:local .
```

## Docker run

```bash
docker run --name go-auth --rm -p 8082:8080 --env-file .env go-auth:local
```


