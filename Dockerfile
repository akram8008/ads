FROM golang:1.18-alpine as build

RUN mkdir /ads

ADD . /ads

WORKDIR /ads

RUN go build -o main ./cmd

FROM alpine:latest
COPY --from=build /ads /ads

WORKDIR /ads

CMD ["/ads/main"]