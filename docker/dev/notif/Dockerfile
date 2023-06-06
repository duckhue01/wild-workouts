# Build the application from source
FROM golang:1.20 AS builder

COPY ./internal/common /internal/common
COPY ./internal/notif /internal/notif

WORKDIR /internal/notif

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/notif

# Move the application to new env
FROM alpine:3.14
WORKDIR /internal

COPY --from=builder /app/notif /internal/service
COPY --from=builder /internal/notif/postgres/migration  /internal/notif/postgres/migration

EXPOSE 3000

CMD [ "./service" ]
