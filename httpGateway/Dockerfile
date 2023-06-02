FROM golang:1.20-buster
ARG DEBIAN_FRONTEND=noninteractive
ARG RUNNING_APP_PORT
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./httpGateway/ ./httpGateway/
COPY ./common/ ./common/
RUN go mod tidy
RUN go install ./httpGateway/cmd/httpGateway/ 

CMD ["httpGateway"]