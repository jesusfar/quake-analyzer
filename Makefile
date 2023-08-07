# Variables
BINARY_NAME=quake-analyzer
SRC_DIR=cmd/cli
BUILD_DIR=build
TEST_DIR=./...

# Targets
all: clean build

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)/*.go

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

test:
	@echo "Running tests..."
	@go test -v $(TEST_DIR)

run:
	@echo "Running..."
	@go run $(SRC_DIR)/*.go
