# Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
# Genera go.sum si no existe (p. ej. en entornos sin red local)
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /api ./cmd/api

# Run
FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /api .
EXPOSE 8080
CMD ["./api"]
