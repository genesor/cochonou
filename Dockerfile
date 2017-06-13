# build stage
FROM golang:1.8.3-alpine AS build-env

RUN apk --no-cache add ca-certificates && update-ca-certificates

ADD . /go/src/github.com/genesor/cochonou
WORKDIR /go/src/github.com/genesor/cochonou
RUN  go install -v ./cmd/cochonou/

# final stage
FROM alpine:3.6

RUN apk --no-cache add ca-certificates && update-ca-certificates

EXPOSE 9494
EXPOSE 9393

WORKDIR /app
COPY --from=build-env /go/bin/cochonou /app/
COPY --from=build-env /go/src/github.com/genesor/cochonou/web /app/
ENTRYPOINT ./cochonou