protoc -I. --go_out=plugins=grpc:$GOPATH/src ./idl/*.proto
# protoc --proto_path=. --go_out=plugins=grpc:$GOPATH/src ./idl/*.proto 

