# auth-service/Dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod edit -module=auth-service
RUN go mod tidy
COPY . .
RUN go build -o auth-service
CMD ["./auth-service"]