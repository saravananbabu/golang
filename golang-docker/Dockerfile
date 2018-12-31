FROM  golang

ADD . /go/src/github.com/saravananbabu/golang/golang-docker

RUN go install github.com/saravananbabu/golang/golang-docker

ENTRYPOINT /go/bin/golang-docker

EXPOSE 8080