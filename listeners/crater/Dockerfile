FROM golang

#Install git
RUN apt-get update; apt-get install -y git
RUN mkdir -p /go/src/github.com/degenerat3/
RUN cd /go/src/github.com/degenerat3; git clone https://github.com/degenerat3/meteor; cd meteor; git checkout crater-dev;
WORKDIR /go/src/github.com/degenerat3/meteor 
RUN go mod tidy
RUN cd /go/src/github.com/degenerat3/meteor/listeners/crater; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o crater .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/degenerat3/meteor/listeners/crater .
CMD ["./crater"] 