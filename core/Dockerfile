FROM golang

#Install git
RUN apt-get update; apt-get install -y git
RUN mkdir -p /go/src/github.com/degenerat3/
RUN cd /go/src/github.com/degenerat3; git clone https://github.com/degenerat3/meteor; cd meteor; 
RUN go mod tidy
RUN cd /go/src/github.com/degenerat3/meteor/core; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o core .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/degenerat3/meteor/core .
# Add docker-compose-wait tool -------------------
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait
CMD ["./core"] 