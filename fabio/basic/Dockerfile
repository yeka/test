FROM golang:1.12.9-alpine AS builder
COPY main.go .
RUN go build -o /testapp main.go

FROM alpine:latest
WORKDIR /test
COPY --from=builder /testapp .
ENTRYPOINT ["./testapp"]