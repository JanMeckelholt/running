FROM golang:1.20-buster
ARG DEBIAN_FRONTEND=noninteractive
ARG GPG_PRIVATE_KEY
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
#COPY .env.docker.postgres .
#COPY .env.docker.postgres.secret .
RUN go mod download
COPY ./database-service/. ./database-service/
COPY ./common/ ./common/
RUN echo -n "${GPG_PRIVATE_KEY}" | base64 --decode > private.key
RUN gpg --import private.key
RUN gpg -d ./common/.env.docker.postgres.secret.asc > ./common/.env.docker.postgres.secret
RUN go mod tidy
RUN go install ./database-service/cmd/database-service/
RUN chmod +x ./database-service/wait-for-it.sh ./database-service/docker-entrypoint.sh
ENTRYPOINT ["./database-service/docker-entrypoint.sh"]
CMD ["database-service"]
EXPOSE 666