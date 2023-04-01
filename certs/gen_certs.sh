#!/usr/bin/env bash

rm *.pem
rm ../database-service/certs/*.pem
rm ../strava-service/certs/*.pem
rm ../runner/certs/*.pem

# Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=DE/ST=x/L=x/O=x/CN=www.janmeckelholt.de" 
cp ca-cert.pem ../database-service/certs/
cp ca-cert.pem ../strava-service/certs/
cp ca-cert.pem ../runner/certs/

# database-service
# Generate database-service-server private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout database-service-server-key.pem -out database-service-server-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=database-service"
# Use CA's private key to sign database-service-server's CSR and get back the signed certificate
openssl x509 -req -in database-service-server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out database-service-server-cert.pem -extfile ./altNames_database-service.cnf
mv database-service-server-cert.pem ../database-service/certs/
mv database-service-server-key.pem ../database-service/certs/

# strava-service
# Generate strava-service-server private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout strava-service-server-key.pem -out strava-service-server-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=strava-service"
# Use CA's private key to sign strava-service-server's CSR and get back the signed certificate
openssl x509 -req -in strava-service-server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out strava-service-server-cert.pem -extfile ./altNames_strava-service.cnf
mv strava-service-server-cert.pem ../strava-service/certs/
mv strava-service-server-key.pem ../strava-service/certs/

