FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o auth-service .
EXPOSE 8000
CMD ["./auth-service"]
