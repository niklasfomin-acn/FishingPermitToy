FROM golang:1.21-alpine

WORKDIR /AnglerAppKlon

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /AnglerAppKlon/server

EXPOSE 3000

CMD ["/AnglerAppKlon/server"]
