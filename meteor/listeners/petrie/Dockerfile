FROM golang

#Install git
RUN apt-get update; apt-get install -y git
RUN mkdir -p /go/src/github.com/degenerat3/
RUN mkdir /hostedfiles/
RUN cd /go/src/github.com/degenerat3; git clone https://github.com/degenerat3/meteor; cd meteor; git checkout rewrite
RUN go get github.com/degenerat3/meteor/meteor/...
RUN go install github.com/degenerat3/meteor/meteor/listeners/petrie

ENTRYPOINT [ "/go/bin/petrie" ]