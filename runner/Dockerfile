FROM golang:1.20-buster
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
COPY ./runner/ /asset
COPY ./common/ /asset
RUN go mod download
COPY ./runner/ .
RUN go mod tidy
RUN pwd
RUN ls -a
RUN go install ./cmd/runner/
RUN chmod +x ./wait-for-it.sh ./docker-entrypoint.sh
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["runner"]
EXPOSE 80