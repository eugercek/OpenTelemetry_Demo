FROM golang:1.19-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY user/* .

RUN go build -o /user

EXPOSE 8080

CMD [ "/user" ]
