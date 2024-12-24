FROM golang:1.23.0 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y fuse3

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM alpine:latest

COPY --from=builder /main /main

RUN apk add --no-cache fuse3

EXPOSE 8080

RUN mkdir flubber-fuse/

CMD ["/main", "mount", "flubber-fuse"]
