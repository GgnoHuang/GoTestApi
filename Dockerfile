# build
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app

# run
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /src/app .
EXPOSE 8080
CMD ["./app"]
