FROM golang:1.20-buster
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
COPY .env.docker.postgres .
COPY .env.docker.postgres.secret .
RUN go mod download
COPY ./database-service/. ./database-service/
COPY ./common/ ./common/
RUN go mod tidy
RUN go install ./database-service/cmd/database-service/
RUN chmod +x ./database-service/wait-for-it.sh ./database-service/docker-entrypoint.sh
ENTRYPOINT ["./database-service/docker-entrypoint.sh"]
CMD ["database-service"]
EXPOSE 666