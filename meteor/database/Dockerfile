FROM library/postgres:latest

# Have DB/user get created on startup
COPY ini.sql /docker-entrypoint-initdb.d/

EXPOSE 5432

CMD [ "postgres"]