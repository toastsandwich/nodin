build:
	@go build -o nodin

run:
	@go run ./abstract_syntax_tree.go ./compiler.go ./lexer.go ./main.go ./visitor.go

store:
	sudo mv nodin /usr/local/bin/

install: build store
