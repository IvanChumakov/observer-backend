FROM golang:alpine AS builder

WORKDIR /observer-backend
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY docs ./docs
RUN CGO_ENABLED=0 GOOS=linux go build -o observer ./cmd/observer/main.go

CMD ["./observer"]