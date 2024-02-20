BINARY_NAME=gowc
INSTALL_DIR=/usr/local/bin

build:
	go build -o $(BINARY_NAME)

install:
	@make build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)."
	@sudo cp $(BINARY_NAME) $(INSTALL_DIR)
	@chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installation complete."

run: 
	make install
	@echo "\n"
	@gowc

