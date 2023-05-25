FROM ubuntu/postgres:14-22.04_beta
ARG DEBIAN_FRONTEND=noninteractive
ARG GPG_PRIVATE_KEY
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /
COPY ./postgres/ /postgres/
COPY ./common/ /common/
RUN echo -n "${GPG_PRIVATE_KEY}" | base64 --decode > private.key && \
gpg --import private.key && \
gpg -d ./common/.env.docker.postgres.secret.asc > ./common/.env.docker.postgres.secret && \
gpg -d ./postgres/certs/postgres-key.pem.asc > ./postgres/certs/postgres-key.pem && \
rm -f private.key && unset GPG_PRIVATE_KEY && \
export $( grep -vE "^(#.*|\s*)$" ./common/.env.docker.postgres.secret) && \
export $( grep -vE "^(#.*|\s*)$" ./common/.env.docker.postgres)
RUN chown root:1001 ./postgres/certs/postgres-key.pem
RUN chmod 640 ./postgres/certs/postgres-key.pem

ENTRYPOINT [ "postgres", "-c", "ssl=on", "-c", "ssl_cert_file=/postgres/certs/postgres-cert.pem", "-c", "ssl_key_file=/postgres/certs/postgres-key.pem"]
EXPOSE 5432
