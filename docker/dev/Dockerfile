# Build the application from source
FROM golang:1.20 AS builder

ARG SERVICE

COPY ./internal/common /internal/common
COPY ./internal/$SERVICE /internal/$SERVICE

WORKDIR /internal/$SERVICE

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/$SERVICE

# Move the application to new env
FROM alpine:3.14
ARG SERVICE

COPY --from=builder /app/$SERVICE /internal/$SERVICE

EXPOSE 3000

CMD [ "/internal/${SERVICE}" ]
