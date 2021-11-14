appName = webapp

# Run application in development mode.
dev.run:
	go run ./cmd/$(appName)/. -devMode

# Run application.
run:
	go run ./cmd/$(appName)/.

# Build application.
build:
	@echo "Starting the application build process..."
	go build -o bin/$(appName) ./cmd/$(appName)/.
	@echo Finished.
	@echo bin/$(appName)
