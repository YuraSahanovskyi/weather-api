FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/weather-api

# Final image
FROM alpine

WORKDIR /

COPY --from=build /app/app .

ENTRYPOINT ["/app"]
