FROM golang

COPY . /go/src/github.com/degenerat3/meteor/meteor/listeners/daddy_tops
WORKDIR /app
RUN go get ./...
RUN go build

ENTRYPOINT [ "/app/daddy_tops" ]