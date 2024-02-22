# Makefile

PROGRAM_NAME = fancyclock
SRC_DIR := $(shell pwd)/src/main.go
ASSETS_DIR = $(shell pwd)/src/assets
SRC_FILES = $(wildcard $(shell pwd)/src/*.go)
ASSETS_FILES = $(wildcard $(ASSETS_DIR)/*)

BIN_DIR = $(GOPATH)/bin

INSTALL_PATH = $(BIN_DIR)/$(PROGRAM_NAME)

all: $(PROGRAM_NAME)

$(PROGRAM_NAME): $(SRC_FILES)
	go build -o $(PROGRAM_NAME) $(SRC_DIR)

install: $(PROGRAM_NAME)
	@mkdir -p $(BIN_DIR)
	@sudo cp $(PROGRAM_NAME) $(BIN_DIR)
	@sudo cp -r $(ASSETS_DIR) $(BIN_DIR)

uninstall:
	@sudo rm -f $(INSTALL_PATH)
	@sudo rm -rf $(BIN_DIR)/assets

clean:
	@rm -f $(PROGRAM_NAME)

.PHONY: all install uninstall clean
