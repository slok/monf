FROM golang:1.5.1
MAINTAINER Xabier Larrakoetxea <slok69@gmail.com>

# Create the user/group for the running stuff
RUN groupadd -g 1000 monf
RUN useradd -m -u 1000 -g 1000 monf
RUN chown monf:monf -R /go
#RUN chmod 777 /usr/src/go/pkg/tool/linux_amd64/
USER monf

# Install handy dependencies/tools
RUN go get github.com/tools/godep
RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/axw/gocov/gocov
RUN go get github.com/mailgun/godebug

# Install goconvey for testing
RUN go get github.com/smartystreets/goconvey

# Install project dependencies
COPY . /tmp/monf
RUN cd /tmp/monf && godep restore

WORKDIR /go/src/github.com/slok/monf
