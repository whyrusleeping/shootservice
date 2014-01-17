all: client server

server:
	go build server.go

client:
	go build client.go
