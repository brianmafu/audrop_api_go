# ======================================================================================================================
#       Use an official Golang runtime as a Parent Image
# ======================================================================================================================

FROM golang:alpine as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .

FROM alpine

ENV DB_NAME=mymtnshop
ENV DB_HOST="mongodb_mongodb_1:27017"
ENV DB_USERNAME=""
ENV DB_PASSWORD=""

COPY --from=builder /build/main /app/

WORKDIR /app

CMD ["./main"]
