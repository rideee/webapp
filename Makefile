build:
	@echo "Starting the application build process..."
	go build -o bin/webapp ./cmd/webapp/.
	@echo -e "\nCheck the bin folder."

run:
	go run ./cmd/webapp/.
