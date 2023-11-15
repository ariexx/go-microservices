FROM alpine:latest

RUN mkdir /app

COPY orderApp /app

CMD ["/app/orderApp"]