FROM alpine:latest
RUN apk add --update tzdata

ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY ./mctest_serv /bin/mctest_server
RUN ["chmod", "+x", "/bin/mctest_server"]
CMD ["/bin/mctest_server"]