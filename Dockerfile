FROM golang:1.16.0-stretch AS builder
WORKDIR /app
COPY . .
RUN make compile

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/tetesan-hujan .
CMD ["/app/tetesan-hujan"]