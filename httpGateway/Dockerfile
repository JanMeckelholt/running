FROM golang:1.20-buster
ARG DEBIAN_FRONTEND=noninteractive
ARG GPG_PRIVATE_KEY
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./httpGateway/ ./httpGateway/
COPY ./common/ ./common/

#RUN mkdir -p ~/.gnupg/
#RUN echo $(whoami)
#RUN chown -R $(whoami) ~/.gnupg/
RUN echo -n "${GPG_PRIVATE_KEY}" | base64 --decode > private.key
RUN gpg --import private.key
RUN gpg -d ./httpGateway/.env.docker.secret.asc > ./httpGateway/.env.docker.secret
RUN gpg -d ./common/.env.docker.running_app.secret.asc > ./common/.env.docker.running_app.secret
RUN go mod tidy
RUN go install ./httpGateway/cmd/httpGateway/
CMD ["httpGateway"]