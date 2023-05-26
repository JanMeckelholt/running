FROM golang:1.20-buster as builder
ARG DEBIAN_FRONTEND=noninteractive
ARG GPG_PRIVATE_KEY
ARG RUNNING_APP_PORT
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./httpGateway/ ./httpGateway/
COPY ./common/ ./common/
RUN echo -n "${GPG_PRIVATE_KEY}" | base64 --decode > private.key && \
gpg --import private.key && \
gpg -d ./httpGateway/.env.docker.secret.asc > ./httpGateway/.env.docker.secret && \
gpg -d ./common/.env.docker.running_app.secret.asc > ./common/.env.docker.running_app.secret && \
gpg -d ./httpGateway/certs/http_gateway-server-key.pem.asc > ./httpGateway/certs/http_gateway-server-key.pem && \
rm -f private.key && unset GPG_PRIVATE_KEY
RUN go mod tidy
RUN RUNNING_APP_PORT=${RUNNING_APP_PORT} go install ./httpGateway/cmd/httpGateway/


FROM ubuntu
WORKDIR /httpGateway
COPY --from=builder /go/bin/httpGateway .
CMD ["./httpGateway"]