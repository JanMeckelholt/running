FROM golang:1.20-buster
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
COPY .env.docker.postgres .
COPY .env.docker.postgres.secret .
RUN go mod download
COPY ./token-service/. ./token-service/
COPY ./common/ ./common/
RUN go mod tidy
RUN go install ./token-service/cmd/token-service/
RUN chmod +x ./token-service/wait-for-it.sh ./token-service/docker-entrypoint.sh
ENTRYPOINT ["./token-service/docker-entrypoint.sh"]
CMD ["token-service"]
EXPOSE 666