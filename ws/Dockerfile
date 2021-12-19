FROM golang:1.17-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY server/*.go ./

RUN go build -o /ws-server
EXPOSE 9501

CMD ["/ws-server"]
