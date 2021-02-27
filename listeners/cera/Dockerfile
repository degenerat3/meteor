FROM golang
ARG CORE_SERVER
ARG LOGPATH
ENV corenv=$CORE_SERVER
ENV logenv=$LOGPATH
#Install git
RUN apt-get update; apt-get install -y git
RUN mkdir -p /go/src/github.com/degenerat3/
RUN cd /go/src/github.com/degenerat3; git clone https://github.com/degenerat3/meteor; cd meteor; git checkout crater-dev;
WORKDIR /go/src/github.com/degenerat3/meteor
RUN go mod tidy
RUN mkdir /app
ENTRYPOINT [ "/bin/sh", "-c", "cd /go/src/github.com/degenerat3/meteor/listeners/cera; CGO_ENABLED=0 GOOS=linux go build -ldflags=\"-X 'main.CORESERVER=${corenv}' -X 'main.LOGPATH=${logenv}'\" -a -installsuffix cgo -o cera .; cp cera /app/cera_listener; echo 'CERA LISTENER BUILT - written to ./bin'"]