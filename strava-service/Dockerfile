FROM golang:1.20-buster
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./strava-service/. ./strava-service/
COPY ./common/ ./common/
RUN go mod tidy
RUN go install ./strava-service/cmd/strava-service/

CMD ["strava-service"]
EXPOSE 666