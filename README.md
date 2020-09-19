# Golang REST API Server

## Abstract
- REST API Server
- Language : Go
- Web Framework : Echo (https://github.com/labstack/echo)
- Relational Database : MySQL
- ORM : GORM (https://github.com/jinzhu/gorm)
- Architecture : Clean Architecture

## Explanation Entry
[Go(Echo)×MySQL×Docker×Clean ArchitectureでAPIサーバー構築してみた](https://qiita.com/Le0tk0k/items/c2945c260f28f7ee2d47)

## How To Run This API Server

```
$ git pull https://github.com/Le0tk0k/go-rest-api
$ cd go-rest-api
$ docker-compose up -d
```

## API Docments

**Response example**

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

Get a user by id

```
GET /users/:id
```

Create a user

```
POST /users
```

Update a user

```
PUT /users/:id
```

Delete a user

```
DELETE /users/:id
```

