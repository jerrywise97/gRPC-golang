gen:
	protoc --go-grpc_out=. *.proto
	protoc --go_out=. *.proto



clean:
	rm ./coffee/*.go