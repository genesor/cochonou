NAME            := cochonou
PACKAGES        := $(shell glide novendor)

deps:
	glide up

install:
	glide install

test:
	 go test -v -cover $(PACKAGES)