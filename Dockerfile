# --- Base Stage ---
FROM golang:1.25-alpine AS base
ENV PATH="/go/bin:${PATH}"
RUN apk add --no-cache git
WORKDIR /app

RUN go install github.com/air-verse/air@latest

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod .
RUN go mod download

# --- Development Stage ---
FROM base AS dev
COPY . .

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]
