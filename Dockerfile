FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/api

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
