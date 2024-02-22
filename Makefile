# Makefile

PROGRAM_NAME = fancyclock
SRC_DIR := $(shell pwd)/src/main.go

BIN_DIR = /usr/bin

INSTALL_PATH = $(BIN_DIR)/$(PROGRAM_NAME)

all: $(PROGRAM_NAME)

$(PROGRAM_NAME): $(SRC_FILES)
	go build -o $(PROGRAM_NAME) $(SRC_DIR)

install: $(PROGRAM_NAME)
	@sudo cp $(PROGRAM_NAME) $(BIN_DIR)

uninstall:
	@sudo rm -f $(INSTALL_PATH)

clean:
	@rm -f $(PROGRAM_NAME)

.PHONY: all install uninstall clean
