# REST API with Go and MongoDB
REST API with Golang and MongoDB.
* HTTP router
* CRUD operations
* MongoDB supported driver for Go 
* ENV configurations
* Formatted logs

## Dependencies
``go get github.com/gorilla/mux``  
``go get go.mongodb.org/mongo-driver/mongo``  
``go get github.com/spf13/viper``  
``go get github.com/githubnemo/CompileDaemon``

## Build
``go build``

## Run
``ENV=dev ./go-rest-mongodb``  
``ENV=prod ./go-rest-mongodb``

## Watch
``CompileDaemon -command="./go-rest-mongodb"``

## API
### Postman collection
[Postman collection](/postman-collection.json)

## Resources
mux: A powerful HTTP router and URL matcher for building Go web servers. https://github.com/gorilla/mux

mongo-driver: The MongoDB supported driver for Go. https://github.com/mongodb/mongo-go-driver

viper: Viper is a complete configuration solution for Go applications including 12-Factor apps. https://github.com/spf13/viper

CompileDaemon: Watches your .go files in a directory and invokes go build if a file changed. https://github.com/githubnemo/CompileDaemon
