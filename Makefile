.PHONY: build clean run install uninstall help

PREFIX ?= /usr/local
BINDIR = $(PREFIX)/bin

# Default target
build:
	go build -o gol .

# Clean built binaries
clean:
	rm -f gol

# Build and run with default settings
run: build
	./gol

# Install to system
install: build
	install -d $(BINDIR)
	install -m 755 gol $(BINDIR)

# Remove from system
uninstall:
	rm -f $(BINDIR)/gol

# Show help with available targets
help:
	@echo "Available targets:"
	@echo "  build     - Build the gol binary"
	@echo "  clean     - Remove built binaries"
	@echo "  run       - Build and run with default settings"
	@echo "  install   - Install gol to $(BINDIR)"
	@echo "  uninstall - Remove gol from $(BINDIR)"
	@echo "  help      - Show this help message"