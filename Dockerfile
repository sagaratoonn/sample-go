FROM golang
EXPOSE 8080
WORKDIR /go/src
RUN go get github.com/smartystreets/goconvey
COPY . .
# CMD ["goconvey",  "-host", "0.0.0.0"]
RUN goconvey -host 0.0.0.0 -launchBrowser=false
