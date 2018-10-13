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

___

## Reference commands:

Running micro-api docker container: 

```bash
sudo docker run --network host -p 8080:8080 microhq/micro api --handler=rpc --address=:8080 --namespace=shippy
```
Create User:

```bash
curl -XPOST -H 'Content-Type: application/json' curl -XPOST -H 'Content-Type: application/json' -d '{ "service": "shippy.auth", "method": "Auth.Create", "request":  { "name": "Paul Mauviel", "company": "mauVILLE Technologies", "email": "your@email.com", "password": "SomePass" } }' http://localhost:8080/rpc
```

Curl command to create consignment:

```bash
curl -XPOST -H 'Content-Type: application/json' -H 'Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiNGE4YTk0MTYtY2FlZS00NTczLWFjNGEtMzA0MmE3MGMxMmI2IiwibmFtZSI6IlBhdWwgTWF1dmllbCIsImNvbXBhbnkiOiJtYXVWSUxMRSBUZWNobm9sb2dpZXMiLCJlbWFpbCI6InlvdXJAZW1haWwuY29tIiwicGFzc3dvcmQiOiIkMmEkMTAkT0llOE5Ucm5uMHdUUDZZREVDYUtHdUFZbno0US56U1guUnRGbU5ZSXFBOGF3cjNZVXBZdHUifSwiZXhwIjoxNTM5NjY5MzIzLCJpc3MiOiJtaWNyb3NlcnZpY2VzLnVzZXIifQ.F7d1vgAWsr7NqPjYEtoZF7KR-8Q8xEdfF_LyFVWZX0w' -d '{
      "service": "shippy.consignment",
      "method": "ShippingService.CreateConsignment",
      "request": {
        "description": "This is a test",
        "weight": 500,
        "containers": []
      }
    }' --url http://localhost:8080/rpc
```