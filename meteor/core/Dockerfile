FROM golang

COPY . /go/src/github.com/degenerat3/meteor/meteor/core
#Install git
RUN apt-get update \        
    apt-get install -y git
RUN mkdir -p /go/src/github.com/degenerat3/
RUN cd /go/src/github.com/degenerat3; git clone https://github.com/degenerat3/meteor 
RUN go get github.com/degenerat3/meteor/meteor/...
RUN go install github.com/degenerat3/meteor/meteor/core

ENTRYPOINT [ "/go/bin/core" ]