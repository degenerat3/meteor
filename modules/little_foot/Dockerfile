FROM alpine:latest
RUN apk add --update tzdata

ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY ./lf_server /bin/lf_serv
RUN ["chmod", "+x", "/bin/lf_serv"]
CMD ["/bin/lf_serv"]