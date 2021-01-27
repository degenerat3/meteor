FROM golang

#Install git
RUN apt-get update; apt-get install -y git
RUN mkdir -p /go/src/github.com/degenerat3/
RUN mkdir -p /hostedfiles/dt/
RUN cd /go/src/github.com/degenerat3; git clone https://github.com/degenerat3/meteor; cd meteor; git checkout rewrite
RUN go get github.com/degenerat3/meteor/...
RUN go install github.com/degenerat3/meteor/listeners/nest

ENTRYPOINT [ "/go/bin/nest" ]