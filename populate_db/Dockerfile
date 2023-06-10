FROM golang:1.20-buster
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./populate_db/ ./populate_db/
COPY ./common/ ./common/
RUN go mod tidy
RUN go install ./populate_db/cmd/populate_db/

RUN chmod +x ./populate_db/wait-for-it.sh ./populate_db/docker-entrypoint.sh

ENTRYPOINT ["./populate_db/docker-entrypoint.sh"]
CMD ["populate_db"]
EXPOSE 666