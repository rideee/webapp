appName = webapp

# Run application command.
runApp = go run ./cmd/$(appName)/.

# Run application.
run:
	$(runApp)

# Run application in development mode.
run-dev:
	$(runApp) -dev

# Run application in development mode and open it in default web browser.
run-dev-open:
	$(runApp) -dev -open

# Build application.
build:
	@echo "Starting the application build process..."
	go build -o bin/$(appName) ./cmd/$(appName)/.
	@echo Finished.
	@echo bin/$(appName)
