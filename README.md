# goboil-hex
Golang boilerplate with hexagonal architecture
- PORT : 8081
- PATH : /

## Installation

``` bash
# clone the repo
$ git clone 

# go into app's directory
$ cd my-project
```

## Build & Run

Local
``` bash
# build 
$ go build

# run
$ go run cmd/main.go

```

Docker
``` bash
# build 
$ docker build -t goboil-hex-api:latest .

# config
sudo sysctl -w vm.max_map_count=262144 # it is required for elasticsearch config

# run
$ docker compose up
```

## Documentation

Install environment
``` bash
# get swagger package 
$ go get github.com/swaggo/swag

# move to swagger dir
$ cd $GOPATH/src/github.com/swaggo/swag

# install swagger cmd 
$ go install cmd/swag
```

Generate documentation
``` bash
# generate swagger doc
$ swag init --propertyStrategy snakecase
```
to see the results, run app and access {{url}}/swagger/index.html

# Author
Nakoding Community Team