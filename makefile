build:
	@go build -o nodin

run:
	@go run *.go

store:
	sudo mv nodin /usr/local/bin/

install: build store
