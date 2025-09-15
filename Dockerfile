# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api

# Stage 2: Run
FROM alpine:3.18

WORKDIR /

COPY --from=builder /api /api

EXPOSE 8080

CMD ["/api"]
