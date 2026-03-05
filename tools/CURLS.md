# CURLS

Health check

```bash
curl http://localhost:8080/auth/health
```


Create User

```bash
curl -X POST http://localhost:8080/auth/create \
-H "Content-Type: application/json" \
-d '{"email": "test+1@gmail.com",
    "password": "",
    "full_name": "",
    "name"": "",
    "last_name"": "",
    "nickname"": "",
    "id_card"": "",}' 2> /dev/null
```
