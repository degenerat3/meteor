FROM golang
ARG CORE_SERVER
ARG LOGPATH
#Install git
RUN apt-get update; apt-get install -y git
RUN mkdir -p /go/src/github.com/degenerat3/
RUN cd /go/src/github.com/degenerat3; git clone https://github.com/degenerat3/meteor;
WORKDIR /go/src/github.com/degenerat3/meteor 
RUN go mod tidy
RUN cd /go/src/github.com/degenerat3/meteor/listeners/cera; CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.CORE_SERVER=${CORE_SERVER}' -X 'main.LOGPATH=${LOGPATH}'" -a -installsuffix cgo -o cera .
RUN mkdir /app
COPY /go/src/github.com/degenerat3/meteor/listeners/cera/cera /app/cera_listener
CMD [ "echo" ]