FROM alpine:3.21 AS builder

ARG GOLANG_VERSION=1.23.9

RUN apk update && \
    apk add --no-cache make gcc openssh bash musl-dev openssl-dev ca-certificates nodejs yarn && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

RUN wget https://go.dev/dl/go$GOLANG_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GOLANG_VERSION.linux-amd64.tar.gz && \
    rm go$GOLANG_VERSION.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

ENV PATH /usr/src/node_modules/.bin:$PATH

RUN mkdir /usr/src

RUN mkdir /usr/src/ipmanager

WORKDIR /usr/src/ipmanager

COPY . ./

RUN go mod tidy

RUN cd web && yarn install

RUN make build

FROM alpine:3.21

COPY --from=builder /usr/src/ipmanager/build /usr/bin

RUN mkdir /etc/ipmanager

EXPOSE 8000

CMD ["sh"]
