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