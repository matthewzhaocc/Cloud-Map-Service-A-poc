FROM golang:latest AS builder
WORKDIR /build
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app . 

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/app .
CMD ["./app"]