<!--  -->
1. create go projrct
=> go mod init projrct-name
<!--  -->
<!-- grpc package -->
go get google.golang.org/grpc
go get google.golang.org/protobuf

<!-- grpc cmd -->
protoc --go_out=. --go-grpc_out=. api/grpc/user.proto
