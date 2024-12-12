FROM golang:1.23.0 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y fuse3

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM alpine:latest

# all your keys belong to us
ENV BUCKET_URL=
ENV ACCESS_KEY_ID=
ENV ACCESS_KEY_ID_FILE=
ENV SECRET_ACCESS_KEY=
ENV SECRET_ACCESS_KEY_FILE=
ENV AUTHFILE=
ENV BUCKET_NAME=

COPY --from=builder /main /main

RUN apk add --no-cache fuse3

EXPOSE 8080

RUN mkdir flubber-fuse/

CMD ["/main", "mount", "flubber-fuse"]
