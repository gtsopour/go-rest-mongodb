# REST API with Go and MongoDB
REST API with Go and MongoDB

## Dependencies
``go get github.com/gorilla/mux``  
``go get gopkg.in/mgo.v2``  
``go get github.com/BurntSushi/toml``  
``go get github.com/githubnemo/CompileDaemon``

toml: TOML aims to be a minimal configuration file format that's easy to read due to obvious semantics. https://github.com/toml-lang/toml

mux: A powerful HTTP router and URL matcher for building Go web servers. https://github.com/gorilla/mux

mgo: MongoDB driver

CompileDaemon: Watches your .go files in a directory and invokes go build if a file changed. https://github.com/githubnemo/CompileDaemon

## Watch/Build
``CompileDaemon -command="./go-rest-mongodb"``  
