FROM ubuntu/postgres:14-22.04_beta as builder
ARG DEBIAN_FRONTEND=noninteractive
ARG POSTGRES_PASSWORD
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /
VOLUME /postgres/
COPY ./postgres/ /postgres/

# FIX ME!!!
#COPY ./volumes-data/common /postgres/commonenv/
#COPY ./volumes-data/postgres/certs /postgres/certs/


#RUN export $( grep -vE "^(#.*|\s*)$" /postgres/commonenv/.env.docker.postgres)
#RUN chown root:$USER /postgres/certs/postgres-key.pem
#RUN chmod 640 /postgres/certs/postgres-key.pem


ENTRYPOINT [ "postgres", "-c", "ssl=on", "-c", "ssl_cert_file=/postgres/certs/postgres-cert.pem", "-c", "ssl_key_file=/postgres/certs/postgres-key.pem"]
EXPOSE 5432






