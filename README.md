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

### Article
**Response**

```
[
    {
        "id": 1,
        "title": "article1",
        "url": "article1.com"
    }
]
```

**Endpoint**

Get all articles

```
GET /articles
```

Get an article by id

```
GET /articles/:id
```

Create an article

```
POST /articles
```

Update an article

```
PUT /articles/:id
```

Delete an article

```
DELETE /articles/:id
```

### Category
**Response**

```
{
    "name": "category1"
}
```

**Endpoint**

Get all categories

```
GET /categories
```

Get an category by id

```
GET /categories/:id
```

Create an category

```
POST /categories
```

Update an category

```
PUT /categories/:id
```

Delete an category

```
DELETE /categories/:id
```
