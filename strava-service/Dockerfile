FROM golang:1.20-buster
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
COPY ./strava-service/. /asset
COPY ./common/. /asset
RUN go mod download
COPY ./strava-service/. .
RUN go mod tidy
RUN pwd
RUN ls -a
RUN go env
RUN go install ./cmd/strava-service/
CMD ["strava-service"]
EXPOSE 666