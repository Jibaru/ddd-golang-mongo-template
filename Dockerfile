FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o main cmd/main/main.go

EXPOSE 8080

CMD ["./main"]
