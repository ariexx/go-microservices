#FROM golang:1.20-alpine AS builder
#
#RUN mkdir /app
#
#COPY . /app
#
#WORKDIR /app
#
#RUN #CGO_ENABLED=0 go build -o broker-service ./cmd/api
#RUN CGO_ENABLED=0 go build -o product-client ./cmd/client
#
#RUN #chmod +x ./broker-service
#RUN chmod +x ./product-client
#
##build tiny image
#FROM alpine:latest
#
#RUN mkdir /app
#
##COPY --from=builder /app/broker-service /app
#COPY --from=builder /app/product-client /app
#
#CMD ["/app/product-client"]


FROM alpine:latest

RUN mkdir /app

COPY productClient /app

CMD ["/app/productClient"]