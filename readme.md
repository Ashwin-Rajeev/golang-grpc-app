# Golang gRPC Application with SQLite
A simple gRPC application to query key/value from SQLite database table.

## steps to execute
#### 1. Install the protocol compiler plugins for Go using the following commands:
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```
#### 2.Update your PATH so that the protoc compiler can find the plugins:
```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```
#### 3. start grpc server
```bash
go run grpc-server/main.go
```
#### 4. start grpc client
```bash
go  run grpc-client/main.go
```
## Sample Output 
```
2021/07/03 22:52:25 ******************* Student details *******************
2021/07/03 22:52:25     Student ID      : 1
2021/07/03 22:52:25     Student Name    : Liana Kim
2021/07/03 22:52:25     Student Code    : 0001
2021/07/03 22:52:25     Student Program : Bachelor
```
