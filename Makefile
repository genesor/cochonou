NAME            := cochonou
PACKAGES        := $(shell glide novendor)

deps:
	glide up

install:
	glide install

run:
	go run cmd/cochonou/main.go

test:
	 go test -v -cover $(PACKAGES)