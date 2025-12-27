#1.bulid stage
FROM golang:1.24.4-alpine AS builder

LABEL maintainer="Tadros_Selim"
LABEL discribtion="ASCII Art Transmation Dockerized"
LABEL version="1.0"

WORKDIR /app

COPY . .

RUN go build -o server

#2. RUN STAGE

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY templates ./templates
COPY banners ./banners
COPY static ./static

EXPOSE 8080

CMD ["./server"]
