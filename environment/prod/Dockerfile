FROM golang:1.24.0-alpine3.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api


FROM gcr.io/distroless/static-debian12

WORKDIR /app
COPY --from=builder /app/main .
CMD ["/app/main"]
