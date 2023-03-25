FROM golang:1.20-buster
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
COPY ./strava-service/. /asset
RUN go mod download
COPY ./strava-service/. ./strava-service/
COPY ./common/ ./common/
RUN pwd
RUN go mod tidy
RUN go install ./strava-service/cmd/strava-service/
CMD ["strava-service"]
EXPOSE 666