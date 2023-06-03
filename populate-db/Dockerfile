FROM golang:1.20-buster
ARG DEBIAN_FRONTEND=noninteractive
ARG GPG_PRIVATE_KEY
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./populate-db/. ./populate-db/
COPY ./common/ ./common/
RUN echo -n "${GPG_PRIVATE_KEY}" | base64 --decode > private.key
RUN gpg --import private.key
RUN gpg -d ./http_gateway/.env.docker.secret.asc > .env.docker.secret
RUN rm -f private.key && unset GPG_PRIVATE_KEY
RUN go mod tidy
RUN go install ./populate-db/cmd/populate-db/
RUN chmod +x ./populate-db/wait-for-it.sh ./populate-db/docker-entrypoint.sh
ENTRYPOINT ["./populate-db/docker-entrypoint.sh"]
CMD ["populate-db"]
EXPOSE 666