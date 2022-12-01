FROM golang:1.19.0-buster AS builder
WORKDIR /app
COPY . .
RUN make compile

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/tetesan-hujan-bot .
EXPOSE 8080
CMD ["/app/tetesan-hujan-bot"]
