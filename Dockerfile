# Build stage
FROM golang:1.24.2-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# If main.go in ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /my-go-app ./cmd

# Final stage
FROM alpine:3.18
RUN addgroup -S app && adduser -S app -G app
COPY --from=builder /my-go-app /usr/local/bin/my-go-app
USER app
EXPOSE 8080
ENV APP_VERSION=dev
ENTRYPOINT ["/usr/local/bin/my-go-app"]
