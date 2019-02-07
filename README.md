# GoGin API Starter

GoGin is a Golang & Gin API starter template with some built in niceities to help you hit the ground running.
Swagger is built for robust API documentation and testing.

## Installation
```
$ go get github.com/xxphenomxx/GoGin

## How to run

### Required

- Glide (Install dependencies)
- MongoDB
- Redis
```

### Conf

You should modify `conf/app.ini`

```
[mongo]
User = 
Password =
Host = 127.0.0.1
Posr = 27017
Name = testdb

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### Run
```
$ cd $GOPATH/src/GoGin

[glide](https://glide.sh/)
$ glide install

$ go run main.go 
```

Project information and existing API

```
Initializing MongoDB connection...
Successfully connected to MongoDB
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /export/*filepath         --> GoGin/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
[GIN-debug] HEAD   /export/*filepath         --> GoGin/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
[GIN-debug] GET    /upload/images/*filepath  --> GoGin/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
[GIN-debug] HEAD   /upload/images/*filepath  --> GoGin/vendor/github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
[GIN-debug] GET    /auth                     --> GoGin/vendor/github.com/xxphenomxx/GoGin/routers/api.CheckAuth (4 handlers)
[GIN-debug] GET    /swagger/*any             --> GoGin/vendor/github.com/swaggo/gin-swagger.WrapHandler.func1 (4 handlers)
[GIN-debug] POST   /upload                   --> GoGin/vendor/github.com/xxphenomxx/GoGin/routers/api.UploadImage (4 handlers)
[GIN-debug] GET    /api/v1/users             --> GoGin/vendor/github.com/xxphenomxx/GoGin/routers/api/v1.GetUsers (5 handlers)
[info] start http server listening :9000

Listening port is 9000
Actual pid is 4393
```
Swagger doc

Browse to: 127.0.0.1:9000/swagger/index.html
![image](https://i.imgur.com/nip4Xhr.png)

## Features

- RESTful API
- Golang Mgo
- Swagger
- logging
- Jwt-go
- CORS
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis