gen:
	protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go --port 50051

client:
	go run cmd/client/main.go --address 0.0.0.0:50051

test:
	go test -cover -race ./...