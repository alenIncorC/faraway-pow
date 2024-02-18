# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY . ./

RUN \
    cd cmd && \
    go build -o /powtcpserver -mod vendor

EXPOSE 8080

CMD [ "/powtcpserver" ]