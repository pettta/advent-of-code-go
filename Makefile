.PHONY: all build install clean

all: build install

build:
	@echo "Building aoc..."
	@go build -o aoc ./cmd
	@chmod +x aoc
	@echo "Build complete: ./aoc"

install: build
	@echo "Installing aoc to /usr/local/bin..."
	@sudo mv aoc /usr/local/bin/
	@echo "Installation complete. You can now run 'aoc' from anywhere."

clean:
	@echo "Cleaning up..."
	@rm -f aoc
	@echo "Clean complete."

