.PHONY: help
help:
	@echo 'make [command]'
	@echo '	build'

.PHONY: build
build:
	@go build -o youdont ./cmd/*.go
