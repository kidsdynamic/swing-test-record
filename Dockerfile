FROM golang:1.8.3
  RUN mkdir -p /go/src/github.com/swing-test-record/
 ADD ./build /go/src/github.com/swing-test-record/
WORKDIR /go/src/github.com/swing-test-record
 CMD ["/go/src/github.com/swing-test-record/main"]

  EXPOSE 8110

