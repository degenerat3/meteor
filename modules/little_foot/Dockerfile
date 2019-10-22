FROM golang:1.7.3
WORKDIR /code/
RUN go get github.com/degenerat3/metcli
COPY lf_server.go .
RUN GOOS=linux go build -o lf_server .

FROM alpine:latest
RUN apk add --update tzdata

ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=0 /code/lf_server /bin/lf_serv
RUN ["chmod", "+x", "/bin/lf_serv"]
CMD ["/bin/lf_serv"]