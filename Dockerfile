FROM golang

ADD . /go/src/golang-starter

RUN go get -u github.com/pineda89/golang-springboot/...
#RUN go get github.com/Jeffail/gabs
#RUN go get github.com/satori/go.uuid
#RUN go get github.com/hudl/fargo

RUN go install golang-starter

ENTRYPOINT /go/bin/golang-starter

EXPOSE 8080
