FROM ubuntu as builder
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Berlin"
RUN apt-get update
RUN apt-get install -y curl git wget unzip libgconf-2-4 gdb libstdc++6 libglu1-mesa fonts-droid-fallback gnupg2 ca-certificates cl-base64 
RUN apt-get clean
RUN git clone https://github.com/flutter/flutter.git /usr/local/flutter
ENV PATH="/usr/local/flutter/bin:/usr/local/flutter/bin/cache/dart-sdk/bin:${PATH}"
RUN flutter doctor -v
RUN flutter channel master
RUN flutter upgrade
RUN flutter config --enable-web
WORKDIR /

COPY ./running_app/ ./running_app/

# FIX ME!!!
#COPY ./volumes-data/running_app/certs ./running_app/certs
#RUN chown root:$USERID ./running_app/certs/running_app-key.pem
#RUN chmod 640 ./running_app/certs/running_app-key.pem
WORKDIR /running_app/
RUN flutter create .
RUN flutter clean
RUN flutter pub get
RUN flutter pub upgrade --major-versions
RUN flutter build web --dart-define=RUNNING_APP_PASSWORD=${RUNNING_APP_PASSWORD} 


FROM ubuntu
RUN apt-get update && apt-get install -y python3 
WORKDIR /
RUN mkdir -p running_app/certs
RUN mkdir common
COPY --from=builder /running_app/build/web /running_app
COPY --from=builder /running_app/certs /running_app/certs
COPY --from=builder /running_app/server.* /running_app/
RUN chmod +x /running_app/server.sh 

EXPOSE 5000

ENTRYPOINT [ "/running_app/server.sh"]