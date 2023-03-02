FROM alpine:latest as alpine

RUN apk add -U --no-cache ca-certificates


FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN make build


FROM scratch

WORKDIR /app

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/bin/rumgo rumgo

CMD ["/app/rumgo"]
