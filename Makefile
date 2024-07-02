# Makefile content
# Makefile for Go project

# Variables
SCRIPT_PATH := ./scripts/run_dev.sh

# Default target
.PHONY: dev
dev:
	@$(SCRIPT_PATH)