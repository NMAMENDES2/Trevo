APP_NAME := Trevo
CMD_DIR := cmd
BIN_DIR := bin
GO := go

.PHONY: all
all: build

.PHONY: build
build:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

.PHONY: run
run: build
	./$(BIN_DIR)/$(APP_NAME)

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)/$(APP_NAME)

