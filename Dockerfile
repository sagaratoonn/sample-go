FROM golang
WORKDIR /go/src
RUN go get github.com/smartystreets/goconvey
COPY . .
RUN ${GOPATH}/bin/goconvey
