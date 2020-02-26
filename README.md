# REST API with Go and MongoDB
REST API with Go and MongoDB

## Dependencies
``go get github.com/gorilla/mux``  
``go get go.mongodb.org/mongo-driver/mongo``  
``go get github.com/spf13/viper``  
``go get github.com/githubnemo/CompileDaemon``

mux: A powerful HTTP router and URL matcher for building Go web servers. https://github.com/gorilla/mux

mongo-driver: The MongoDB supported driver for Go. https://github.com/mongodb/mongo-go-driver

viper: Viper is a complete configuration solution for Go applications including 12-Factor apps. https://github.com/toml-lang/toml

CompileDaemon: Watches your .go files in a directory and invokes go build if a file changed. https://github.com/spf13/viper

## Watch/Build
``CompileDaemon -command="./go-rest-mongodb"``  
