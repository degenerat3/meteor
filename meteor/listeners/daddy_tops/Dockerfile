FROM golang

COPY . /go/src/github.com/degenerat3/meteor/meteor/listeners/daddy_tops
RUN go get github.com/degenerat3/meteor/meteor/listeners/daddy_tops/...
RUN go install github.com/degenerat3/meteor/meteor/listeners/daddy_tops

ENTRYPOINT [ "/go/bin/daddy_tops" ]