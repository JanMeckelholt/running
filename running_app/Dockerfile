FROM ubuntu:23.04
ARG DEBIAN_FRONTEND=noninteractive
ARG RUNNING_APP_PASSWORD
ENV TZ="Europe/Berlin"
RUN apt-get update
RUN apt-get install -y curl git wget unzip libgconf-2-4 gdb libstdc++6 libglu1-mesa fonts-droid-fallback lib32stdc++6 python3
RUN apt-get clean
RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter
ENV PATH="/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin:${PATH}"
RUN flutter doctor -v
RUN flutter channel master
RUN flutter upgrade
RUN flutter config --enable-web

WORKDIR /running_app/
COPY ./running_app/ .
RUN flutter create .
RUN flutter clean
RUN flutter pub get
RUN flutter pub upgrade --major-versions
RUN flutter build web --dart-define=RUNNING_APP_PASSWORD=${RUNNING_APP_PASSWORD} --web-use-https --web-ssl-cert-path=./certs/running_app-cert.pem --web-ssl-key-path=./certs/running_app-key.pem

EXPOSE 5000

RUN ["chmod", "+x", "/running_app/server.sh"]

ENTRYPOINT [ "/running_app/server.sh"]