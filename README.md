<!--  -->
1. create go projrct
=> go mod init projrct-name

2. graphql
=> go get github.com/99designs/gqlgen

3.
=> go run github.com/99designs/gqlgen generate

4
=> go get github.com/lib/pq


<!--  -->
<!-- grpc package -->
go get google.golang.org/grpc
go get google.golang.org/protobuf

<!-- grpc cmd -->
protoc --go_out=. --go-grpc_out=. api/grpc/user.proto


<!--  -->
go mod tidy