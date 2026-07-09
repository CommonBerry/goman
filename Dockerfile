FROM golang:1.26.4-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/goman ./cmd/api

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /out/goman /app/goman

EXPOSE 3000

CMD ["/app/goman"]
