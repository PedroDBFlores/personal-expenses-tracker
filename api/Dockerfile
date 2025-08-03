# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -tags netgo -o api

# Deploy stage
FROM scratch

WORKDIR /app
COPY --from=builder /app/api .

ENV PORT=8080
ENV DATABASE_URL=expenses.db

EXPOSE 8080

CMD ["./api"]
