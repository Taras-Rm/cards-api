FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o app main.go

CMD [ "./app" ]