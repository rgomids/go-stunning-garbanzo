build:
	@cd command; \
	go build -v -o go-stunning-garbanzo

buildRun: build
	@./command/go-stunning-garbanzo

run:
	@go run command/main.go