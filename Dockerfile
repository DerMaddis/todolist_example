FROM postgres
ARG password
ENV POSTGRES_PASSWORD ${password}
COPY init.sql /docker-entrypoint-initdb.d/
