FROM golang

COPY . /go/src/github.com/degenerat3/meteor/meteor/core
RUN go get github.com/degenerat3/meteor/meteor/core/...
RUN go install github.com/degenerat3/meteor/meteor/core

ENTRYPOINT [ "/go/bin/core" ]