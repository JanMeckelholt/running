FROM golang:1.20-buster
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./runner/ ./runner/
COPY ./common/ ./common/
RUN go mod tidy
RUN go install ./runner/cmd/runner/
RUN chmod +x ./runner/wait-for-it.sh ./runner/docker-entrypoint.sh
ENTRYPOINT ["./runner/docker-entrypoint.sh"]
CMD ["runner"]
EXPOSE 666