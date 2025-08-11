FROM golang:1.23-alpine

WORKDIR /go/src/nomado-houses

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o /nomado-houses

EXPOSE 8080

CMD ["/nomado-houses"]