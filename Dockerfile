FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o go-rate ./

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/go-rate .
EXPOSE 8080
CMD ["./go-rate"]