FROM alpine:latest
RUN apk add --update tzdata

ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY ./spike_server /bin/spike_serv
RUN ["chmod", "+x", "/bin/spike_serv"]
CMD ["/bin/spike_serv"]