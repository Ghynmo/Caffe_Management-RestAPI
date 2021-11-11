FROM golang:1.16-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app/app
RUN go clean --modcache
RUN go build -o main
# EXPOSE 8080
# CMD ["/app/main"]

FROM alpine:3.14
WORKDIR /root
COPY --from=builder /app/app/config_docker.json .
COPY --from=builder /app/app/main .
EXPOSE 8080
CMD [ "./main" ]