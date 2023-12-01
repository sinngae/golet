.PHONY: all servlet

servlet:
	go build -o ./bin/servlet ./cmd/servlet/main.go

