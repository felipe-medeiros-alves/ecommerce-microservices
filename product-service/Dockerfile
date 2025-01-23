FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o product-service .
EXPOSE 8000
CMD ["./product-service"]
