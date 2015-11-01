FROM golang:1.5.1
MAINTAINER Xabier Larrakoetxea <slok69@gmail.com>


# Install node & npm stuff
RUN apt-get update &&\
    apt-get install sudo


# Prepare npm stuff
RUN curl --silent --location https://deb.nodesource.com/setup_5.x | bash -
RUN apt-get install --yes nodejs
RUN mkdir -p /monf/node_deps
ENV PATH="/monf/node_deps/node_modules/.bin:$PATH"

# Create the user/group for the running stuff
RUN groupadd -g 1000 monf
RUN useradd -m -u 1000 -g 1000 monf
RUN chown monf:monf -R /go
RUN chown monf:monf -R /monf

USER monf

# Install handy dependencies/tools
RUN go get github.com/tools/godep
RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/axw/gocov/gocov
RUN go get github.com/mailgun/godebug

# Install goconvey for testing
RUN go get github.com/smartystreets/goconvey


# Install node dependencies
WORKDIR /monf/node_deps
COPY package.json ./package.json
RUN npm install

# Install project dependencies
COPY . /tmp/monf
RUN cd /tmp/monf && godep restore


# Set settings for dev
ENV MONF_SETTINGS /go/src/github.com/slok/monf/environment/dev/settings.yaml
WORKDIR /go/src/github.com/slok/monf
