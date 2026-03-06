# CURLS

Health check

```bash
curl http://localhost:8080/auth/health
```


Create User

```bash
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{"email": "test+1@test.com", "password": "123456Abcde"}' | jq
```

```bash
curl -X POST http://localhost:8080/auth/create \
-H "Content-Type: application/json" \
-d '{"email": "test+1@test.com",
    "password": "123456",
    "full_name": "test 1",
    "name"": "test1",
    "last_name"": "test1",
    "nickname"": "",
    "id_card"": "1234567890"}' 2> /dev/null
```
