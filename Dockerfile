# build stage
FROM golang:alpine AS build-env

RUN apk --update add ca-certificates && rm -rf /var/cache/apk/*

ADD . /go/src/github.com/genesor/cochonou
RUN cd /go/src/github.com/genesor/cochonou && go build ./cmd/cochonou/

# final stage
FROM alpine

WORKDIR /app

EXPOSE 9494

COPY --from=build-env /go/src/github.com/genesor/cochonou/cochonou /app/
ENTRYPOINT ./cochonou