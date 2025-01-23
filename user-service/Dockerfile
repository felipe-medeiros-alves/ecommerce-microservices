FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o user-service .
EXPOSE 8000
CMD ["./user-service"]
