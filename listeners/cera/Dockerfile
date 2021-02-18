FROM golang

#Install git
RUN apt-get update; apt-get install -y git
RUN mkdir -p /go/src/github.com/degenerat3/
RUN cd /go/src/github.com/degenerat3; git clone https://github.com/degenerat3/meteor;
WORKDIR /go/src/github.com/degenerat3/meteor 
RUN go mod tidy
RUN cd /go/src/github.com/degenerat3/meteor/listeners/cera; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cera .
#COPY /go/src/github.com/degenerat3/meteor/listeners/cera/cera /app/cera_listener