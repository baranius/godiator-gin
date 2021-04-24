PROJECTNAME=$(shell basename "$(PWD)")

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

run:
	@echo "$(PROJECTNAME) is running"
	go run ./main.go
.PHONY: run

format:
	go fmt ./internal/...
.PHONY: format

test: format
	@echo "$(PROJECTNAME) tests are running"
	go test ./internal/...
.PHONY: test

test-coverage: format
	@echo "$(PROJECTNAME) test coverages are running"
	rm -rf .build/test-coverage
	mkdir .build/test-coverage
	go test -coverprofile .build/test-coverage/coverage.out ./internal/...
	go tool cover -html=.build/test-coverage/coverage.out -o .build/test-coverage/cover.html
	open .build/test-coverage/cover.html
.PHONY: test-coverage
