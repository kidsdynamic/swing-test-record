FROM golang:1.7

# glide
RUN \
  wget https://github.com/Masterminds/glide/releases/download/v0.11.1/glide-v0.11.1-linux-amd64.tar.gz && \
  tar xvf glide-v0.11.1-linux-amd64.tar.gz && \
  mv linux-amd64/glide /usr/bin/

#gin
RUN go get github.com/codegangsta/gin

WORKDIR /go/src/swing-test-record

EXPOSE 8110