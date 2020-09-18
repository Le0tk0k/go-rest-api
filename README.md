# Golang REST API Server

## Abstract
- REST API Server
- Language : Go
- Web Framework : Echo (https://github.com/labstack/echo)
- Relational Database : MySQL
- ORM : GORM (https://github.com/jinzhu/gorm)
- Architecture : Clean Architecture

## How To Run This API Server

```
$ git pull https://github.com/Le0tk0k/go-rest-api
$ cd go-rest-api
$ docker-compose up -d
```

## API Docments

### User
**Response**

```
[
    {
        "id": 1,
        "name": "user1",
        "age": 20
    }
]
```

**Endpoint**

Get all users

```
GET /users
```

Get an user by id

```
GET /users/:id
```

Create an user

```
POST /users
```

Update an user

```
PUT /users/:id
```

Delete an user

```
DELETE /users/:id
```

