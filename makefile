run:
	@echo "Starting the Server"
	@cd api_server/ && node main.js &
	@cd  cmd/server/ && go run main.go   