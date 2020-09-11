FROM golang

COPY . /go/src/github.com/degenerat3/meteor/meteor/listeners/daddy_tops
WORKDIR /go/src/github.com/degenerat3/meteor/meteor/listeners/daddy_tops
RUN go get ./...
RUN go build

ENTRYPOINT [ "/app/daddy_tops" ]