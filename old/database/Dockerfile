FROM library/postgres:12

# Have DB/user get created on startup
COPY ini.sql /docker-entrypoint-initdb.d/

EXPOSE 5432

CMD [ "postgres"]