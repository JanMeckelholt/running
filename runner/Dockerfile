FROM golang:1.20-buster as builder
ARG DEBIAN_FRONTEND=noninteractive
ARG GPG_PRIVATE_KEY
RUN apt-get -yq update && apt-get -yq install gnupg2 ca-certificates cl-base64
WORKDIR /go/src/running
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./runner/ ./runner/
COPY ./common/ ./common/
RUN echo -n "${GPG_PRIVATE_KEY}" | base64 --decode > private.key && \
gpg --import private.key && \
gpg -d ./runner/certs/runner-server-key.pem.asc > ./runner/certs/runner-server-key.pem && \
rm -f private.key && unset GPG_PRIVATE_KEY
RUN go mod tidy
RUN go install ./runner/cmd/runner/

FROM alpine
WORKDIR /runner
COPY --from=builder /go/bin/runner .
COPY --from=builder /go/src/running/runner/*.sh .
RUN chmod +x wait-for-it.sh docker-entrypoint.sh 
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["./runner"]
EXPOSE 666