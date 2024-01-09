set positional-arguments

alias b := build
alias u := update
alias r := run

default:
  @just --list

@run *arg:
    go run main.go $@

test:
	go test ./...

build:
	go build -o dist/app main.go

fmt:
	go fmt ./...

update:
	go get
	go mod tidy
