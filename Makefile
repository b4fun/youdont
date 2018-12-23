.PHONY: help
help:
	@echo 'make [command]'
	@echo '	build-functions'

.PHONY: build-functions
build-functions:
	@./script/build-function.sh get-reddit-user-post
