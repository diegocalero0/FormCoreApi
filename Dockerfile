FROM golang:1.22.4-alpine
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /app

EXPOSE 3000

CMD ["/app"]