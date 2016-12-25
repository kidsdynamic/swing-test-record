FROM golang:1.7
  RUN mkdir /go/src/swing-test-record
 ADD . /go/src/swing-test-record/ 
WORKDIR /go/src/swing-test-record
 RUN go build -o main .
 CMD ["/go/src/github.com/swing-test-record/main"]

  EXPOSE 8110
