FROM alpine:latest
RUN apk add --update tzdata

ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY ./pet_server /app/pet_server
WORKDIR /app

ENTRYPOINT ["pet_server"]