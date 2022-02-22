.PHONY: client
.PHONY: server

client:
	@echo "Removing the client binary"
	rm -f bin/client
	@echo "Building the client binary"
	go build -o bin/client client/main.go client/commands.go
server:
	@echo "Removing the server binary"
	rm -f bin/server
	@echo "Building the server binary"
	go build -o bin/server server/main.go
clear:
	rm -f bin/*