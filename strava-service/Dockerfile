FROM golang:1.20-buster
ARG DEBIAN_FRONTEND=noninteractive
ARG GPG_PRIVATE_KEY
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./strava-service/. ./strava-service/
COPY ./common/ ./common/
RUN echo -n "${GPG_PRIVATE_KEY}" | base64 --decode > private.key
RUN gpg --import private.key
RUN gpg -d ./strava-service/certs/strava-service-server-key.pem.asc > ./strava-service/certs/strava-service-server-key.pem
RUN go mod tidy
RUN go install ./strava-service/cmd/strava-service/
CMD ["strava-service"]
EXPOSE 666