FROM golang:1.24.5-alpine3.22 AS builder

WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN go build -o main .

FROM alpine:3.16.3
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
EXPOSE 8080

CMD [ "./main" ]