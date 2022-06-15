# Microservices cinema

### Users service
**POST** localhost:8080/v1/new_user

Body:
```json
{
    "login": "serega",
    "password": "123",
    "date_birthday": "1995-01-01"
}
```

**POST** localhost:8080/v1/login

Body:
```json
{
    "login": "serega2",
    "password": "123"
}
```

**POST** localhost:8080/api/get_user

Body:
```
749ae576-36a7-4cbf-9e36-240380765221
```

---

### Movies service

**POST** localhost:5000/v1/new_movie

Body:
```json
{
    "name": "Фильм22",
    "description": "Описание",
    "min_age": 25
}
```

**POST** localhost:5000/v1/get_movie_by_uuid

Body:
```
7f302a26-7b91-4fa1-8bc0-9ceaf808bad3
```

**GET** localhost:5000/v1/get_movies

---

### Recommendation service

**POST** localhost:5001/v1/get_recommendation

Body:
```
46818ec6-4e1c-462d-b00d-1754a406987a
```
