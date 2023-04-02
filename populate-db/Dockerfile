FROM golang:1.20-buster
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./populate-db/. ./populate-db/
COPY ./common/ ./common/
RUN go mod tidy
RUN go install ./populate-db/cmd/populate-db/
RUN chmod +x ./populate-db/wait-for-it.sh ./populate-db/docker-entrypoint.sh
ENTRYPOINT ["./populate-db/docker-entrypoint.sh"]
CMD ["populate-db"]
EXPOSE 666