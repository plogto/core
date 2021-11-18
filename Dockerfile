FROM golang:1.16-alpine AS build

LABEL maintainer="Mohammad Mahdi <favecode@gmail.com>"

WORKDIR /app

COPY . ./

ENV https-proxy http://fodev.org:8118

# Install dependencies
RUN go mod download && \
  # Build the app
  GOOS=linux GOARCH=amd64 go build -o main && \
  # Make the final output executable
  chmod +x ./main

FROM alpine:latest

# Install os packages
RUN apk --no-cache add bash

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]

EXPOSE 8080