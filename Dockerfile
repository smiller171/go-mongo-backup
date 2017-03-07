FROM alpine
MAINTAINER Scott Miller <scott.miller171@gmail.com>

RUN mkdir /lib64 && ln -s /lib/ld-musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

COPY mongodump /bin/mongodump
COPY go-mongo-backup go-mongo-backup

EXPOSE 8080

CMD [ "/go-mongo-backup" ]
