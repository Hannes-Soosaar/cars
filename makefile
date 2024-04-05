run:
	@echo "Starting the Server"
	@cd api_server/ && node main.js &
	@cd  cmd/server/ && go run main.go   
go:
	@cd  cmd/server/ && go run main.go   
node:
	@cd api_server/ && node main.js
build:
	@cd api_server/ && npm install
re:
	@cd api_server/ && build run
.PHONY:
	@cd api_server/ && build run re
