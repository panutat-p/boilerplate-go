FROM golang:1.23.5-alpine3.21 as builder
WORKDIR /opt

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o /app cmd

FROM alpine:3.21
COPY --from=builder /app /app
EXPOSE 8080
CMD ["/app"]
