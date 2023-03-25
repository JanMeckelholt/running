FROM golang:1.20-buster
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
COPY ./runner/ /asset
RUN go mod download
COPY ./runner/ ./runner/
COPY ./common/ ./common/
RUN pwd
RUN go mod tidy
RUN go install ./runner/cmd/runner/
RUN chmod +x ./wait-for-it.sh ./docker-entrypoint.sh
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["runner"]
EXPOSE 80