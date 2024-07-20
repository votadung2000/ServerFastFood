FROM golang:1.20.5-alpine3.17 AS builder

WORKDIR /app

# RUN apk update && apk add git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN go mod init SERVERFASTFOOD

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/.env .

# CMD ["./main"]
ENTRYPOINT [ "./main" ]