.PHONY: help
help:
	@echo 'make [command]'
	@echo '	build-functions'
	@echo '	deploy-function'

.PHONY: build-functions
build-functions:
	@./script/build-function.sh get-reddit-user-post

.PHONY: deploy-function
deploy-function:
	@./script/deploy-function.sh
