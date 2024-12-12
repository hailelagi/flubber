FROM golang:1.23.0 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM alpine:latest

COPY --from=builder /main /main

EXPOSE 8080

CMD ["/main"]
