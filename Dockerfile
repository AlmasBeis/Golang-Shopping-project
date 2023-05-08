FROM golang:1.20.2

WORKDIR /app



COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o main ./cmd/main.go

CMD ["./main"]