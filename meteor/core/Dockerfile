FROM golang

COPY . /app
WORKDIR /app
RUN go get ./...
RUN go build

ENTRYPOINT [ "/app/app" ]