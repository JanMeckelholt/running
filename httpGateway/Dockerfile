FROM golang:1.20-buster
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./httpGateway/ ./httpGateway/
COPY ./common/ ./common/
RUN go mod tidy
RUN go install ./httpGateway/cmd/httpGateway/
CMD ["httpGateway"]