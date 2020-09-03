FROM golang:1.7.3
WORKDIR /code/
RUN go get -d -v github.com/degenerat3/metcli
COPY pet_server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o pet_server .

FROM alpine:latest
RUN apk add --update tzdata

ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=0 /code/pet_server /bin/pet_server
RUN ["chmod", "+x", "/bin/pet_server"]
CMD ["/bin/pet_server"]