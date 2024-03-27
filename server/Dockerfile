FROM golang:1.17-alpine3.14

WORKDIR /AnglerAppKlon

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /AnglerAppKlon

EXPOSE 3000

CMD ["/AnglerAppKlon"]
