FROM ubuntu:23.04 as builder
ARG DEBIAN_FRONTEND=noninteractive
ARG GPG_PRIVATE_KEY
ENV TZ="Europe/Berlin"
RUN apt-get update
RUN apt-get install -y curl git wget unzip libgconf-2-4 gdb libstdc++6 libglu1-mesa fonts-droid-fallback python3 gnupg2 ca-certificates cl-base64
RUN apt-get clean
RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter
ENV PATH="/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin:${PATH}"
RUN flutter doctor -v
RUN flutter channel master
RUN flutter upgrade
RUN flutter config --enable-web
WORKDIR /
COPY ./common/ ./common/
COPY ./running_app/ ./running_app/
RUN echo -n "${GPG_PRIVATE_KEY}" | base64 --decode > private.key && \
gpg --import private.key && \
gpg -d ./running_app/certs/running_app-key.pem.asc > ./running_app/certs/running_app-key.pem && \
gpg -d ./common/.env.docker.running_app.secret.asc > ./common/.env.docker.running_app.secret && \
rm -f private.key && unset GPG_PRIVATE_KEY
RUN chown root:1001 ./running_app/certs/running_app-key.pem
RUN chmod 640 ./running_app/certs/running_app-key.pem
RUN export $( grep -vE "^(#.*|\s*)$" ./common/.env.docker.running_app.secret)
WORKDIR /running_app/
RUN flutter create .
RUN flutter clean
RUN flutter pub get
RUN flutter pub upgrade --major-versions
RUN flutter build web --dart-define=RUNNING_APP_PASSWORD=${RUNNING_APP_PASSWORD} 


FROM alpine
COPY --from=builder /running_app/build/web /running_app/
COPY --from=builder /running_app/server.* /running_app/
RUN chmod +x /running_app/server.sh 

EXPOSE 5000

ENTRYPOINT [ "/running_app/server.sh"]