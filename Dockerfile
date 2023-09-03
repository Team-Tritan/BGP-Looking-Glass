
FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o api
FROM alpine
COPY --from=builder /app/api /app/api
EXPOSE 8080
CMD ["/app/api"]