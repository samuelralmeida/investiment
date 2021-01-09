run:
ifeq (, $(shell which wtc 2>/dev/null))
	@echo "\033[31mYOU NEED TO RUN: 'go get -u github.com/rafaelsq/wtc'\033[m" && false
endif
	@wtc