.PHONY: client
.PHONY: server

client:
	@echo "Removing the client binary"
	rm -f bin/client
	@echo "Building the client binary"
	go build -o bin/client client/main.go client/commands.go client/http.go
server:
	@echo "Removing the server binary"
	rm -f bin/server
	@echo "Building the server binary"
	go build -o bin/server server/server.go server/http.go server/db.go
inclient:
	@echo "Removing the client binary"
	rm -f bin/client
	@echo "Building the client binary"
	go install client/main.go client/commands.go client/http.go
inserver:
	@echo "Removing the server binary"
	rm -f bin/server
	@echo "Building the server binary"
	go install server/server.go server/http.go server/db.go
rserver: server
		./bin/server
clear:
	rm -f bin/*
