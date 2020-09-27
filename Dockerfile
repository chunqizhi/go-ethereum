# Build Geth in a stock Go builder container
FROM golang:1.15-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git
RUN go env -w GOPROXY=https://goproxy.cn,direct

ADD . /go-ethereum
RUN cd /go-ethereum && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/sbin/
COPY ./genesis/docker-entrypoint.sh /usr/local/bin/
COPY ./genesis/genesis.json /genesis/ 
VOLUME /usr/local/bin/
VOLUME /genesis/
VOLUME /data/
WORKDIR /data/

EXPOSE 8545 8546 8547 30303 30303/udp
ENTRYPOINT ["docker-entrypoint.sh"]
