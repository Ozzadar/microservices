# Dependencies

1. Install protobuf [here](https://github.com/google/protobuf/releases)

2. Install grpc 
```bash
    go get -u google.golang.org/grpc
```

3.  Install protogen plugin for go:

```bash
    go get -u github.com/golang/protobuf/protoc-gen-go
```
4. Install go micro protogen plugins: 

```bash
    go get -u github.com/micro/protobuf/{proto,protoc-gen-go}
```

5. Install docker

6. Install docker-compose [here](https://docs.docker.com/compose/install/#install-compose)

# Running the project (development)

Ensure Consul is running locally. Easiest way is to run:

```bash
    sudo docker run --name consul -p 8600:8600 -p 8500:8500 -p 8301:8301 -p 8302:8302 -p 8300:8300 -d consul
```

Run the services:

```
    docker-compose up
```

Run just the cli afterwards:

```
    docker-compose run consignment-cli
```