FROM golang:1.12 AS builder
WORKDIR /module
COPY . /module
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

FROM alpine:3
WORKDIR /root/
COPY --from=builder /module/app /root/
COPY config /root/config
COPY ui /root/ui
CMD ["./app"]