FROM golang:1.19-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY greeter/* .

RUN go build -o /greeter

EXPOSE 8080

CMD [ "/greeter" ]
