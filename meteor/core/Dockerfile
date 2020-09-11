FROM golang

COPY . /go/src/github.com/degenerat3/meteor/meteor/core
WORKDIR /go/src/github.com/degenerat3/meteor/meteor/core
RUN go get ./...
RUN go build

ENTRYPOINT [ "/app/core" ]