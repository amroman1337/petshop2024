FROM golang:1.22

COPY go.mod go.sum ./
RUN go mod download

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/pet_service

EXPOSE 8080

CMD ["./main"]