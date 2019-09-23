build:
	@cd command; \
	go build -v -o go-stunning-garbanzo

br: build
	@./command/go-stunning-garbanzo

run:
	@go run command/main.go

install:
	@cd command; \
	go get