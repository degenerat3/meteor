FROM golang

COPY . /go/src/github.com/degenerat3/meteor/meteor/core
WORKDIR /app
RUN go get ./...
RUN go build

ENTRYPOINT [ "/app/core" ]