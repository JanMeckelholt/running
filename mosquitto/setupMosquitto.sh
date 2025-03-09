#!/bin/bash

mkdir -p /root/mosquitto
cp -r /mnt/fritz_nas/mosquitto/. /root/mosquitto/
chown root:root /root/mosquitto/config/pwfile
chmod 0700 /root/mosquitto/config/pwfile

cd /root/mosquitto
docker compose down --remove-orphans
docker compose up -d
cd ..