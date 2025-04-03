# Build stage
FROM golang:1.22.12-alpine3.20 AS builder
WORKDIR /myapp
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /myapp/app

RUN go build -o ../../main .

# Run stage
FROM alpine:3.20
WORKDIR /myapp
COPY --from=builder /main .
EXPOSE 6080
CMD ["./main"]