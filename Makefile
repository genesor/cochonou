NAME            := cochonou
PACKAGES        := $(shell glide novendor)

deps:
	glide up

install:
	glide install

run: dep-start
	go run cmd/cochonou/main.go

test: dep-start
	go test -v -cover $(PACKAGES)

dep-start:
	docker-compose up -d

dep-stop:
	docker-compose stop