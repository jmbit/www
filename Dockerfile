# This is a workaround to get an older version of hugo that still works with 
# the theme used
#FROM docker.io/library/debian AS hugoer
#RUN APT_FRONTEND=noninteractive apt update -y && apt install hugo -y
#WORKDIR /usr/local/src/www
#COPY . .
#RUN hugo version
#RUN cd hugo && hugo

FROM docker.io/library/golang:alpine AS builder
RUN apk update \
    && apk add --no-cache git #\
#    && apk add --no-cache --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community hugo

WORKDIR $GOPATH/src/www-jmbit-de
COPY . .
COPY --from=hugoer /usr/local/src/www/hugo $GOPATH/src/www-jmbit-de/hugo
RUN go get -d -v
#RUN cd hugo && hugo
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/www

FROM scratch
COPY --from=builder /go/bin/www /go/bin/www
ENTRYPOINT ["/go/bin/www"]
