FROM golang

#Install git
RUN apt-get update; apt-get install -y git
RUN mkdir -p /go/src/github.com/degenerat3/
RUN mkdir -p /hostedfiles/dt/
RUN cd /go/src/github.com/degenerat3; git clone https://github.com/degenerat3/meteor;
WORKDIR /go/src/github.com/degenerat3/meteor
RUN go mod tidy
RUN go install github.com/degenerat3/meteor/listeners/nest

ENTRYPOINT [ "/go/bin/nest" ]