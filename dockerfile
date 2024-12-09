FROM golang:latest AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./main.go

FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

WORKDIR /srv/

COPY --from=builder /app/app .
COPY ./env.prd .env

EXPOSE 80

CMD ["./zoweb"]

