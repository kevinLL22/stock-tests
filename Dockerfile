# ---------- Build stage ----------
FROM golang:1.24.3 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o stock-api ./cmd/api

# ---------- Runtime stage ----------
FROM gcr.io/distroless/base-debian12
WORKDIR /app

# Copiamos binario y migraciones
COPY --from=builder /app/stock-api .
COPY migrations ./migrations

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/app/stock-api"]
