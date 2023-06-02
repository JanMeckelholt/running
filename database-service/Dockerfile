FROM golang:1.20-buster 
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./database-service/. ./database-service/
COPY ./common/ ./common/
RUN go mod tidy
RUN go install ./database-service/cmd/database-service/ 
RUN chmod +x ./database-service/wait-for-it.sh ./database-service/docker-entrypoint.sh 
ENTRYPOINT ["./database-service/docker-entrypoint.sh"]
CMD ["database-service"]
EXPOSE 666