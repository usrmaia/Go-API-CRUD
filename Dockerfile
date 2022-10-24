FROM golang:1.18

WORKDIR /app

COPY main.go ./
COPY go.mod ./
COPY go.sum ./
COPY src ./src/

RUN ls /app
RUN go build -o server
EXPOSE 9090

CMD ["/app/server"]