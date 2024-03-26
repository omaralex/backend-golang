FROM postgres:latest

RUN rm -rf /docker-entrypoint-initdb.d/*

COPY scripts.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_DB=dbpharmacy
ENV POSTGRES_USER=userpharmacy
ENV POSTGRES_PASSWORD=userpharmacy