FROM golang:1.23.0 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /flubber

CMD ["/flubber"]
