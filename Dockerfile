FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /task_app ./cmd/main.go
#RUN go build -o /task_app

EXPOSE 7540

ENV LOG_LEVEL=info
ENV TODO_HOST=0.0.0.0
ENV TODO_PORT=7540
ENV TODO_DBFILE=./scheduler.db
ENV TODO_PASSWORD=secret_pass

CMD ["/task_app"]
