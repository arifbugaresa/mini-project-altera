FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o backend-mini-project

EXPOSE 8080

CMD ./backend-mini-project