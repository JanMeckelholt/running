#!/usr/bin/env bash

rm *.pem
rm ../database-service/certs/*.*
rm ../strava-service/certs/*.*
rm ../volumes-data/postgres/certs/*.*
rm ../runner/certs/*.*
rm ../populate-db/certs/*.*
rm ../httpGateway/certs/*.*
rm ../running_app/certs/*.*
rm ../postgres/certs/*.*

# Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=DE/ST=x/L=x/O=x/CN=www.janmeckelholt.de" 
cp ca-cert.pem ../database-service/certs/
cp ca-cert.pem ../strava-service/certs/
cp ca-cert.pem ../volumes-data/postgres/certs/root.crt
cp ca-cert.pem ../runner/certs/
cp ca-cert.pem ../populate-db/certs/
cp ca-cert.pem ../httpGateway/certs/
cp ca-cert.pem ../running_app/certs/
cp ca-cert.pem ../postgres/certs/

# database-service
# Generate database-service-server private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout database-service-server-key.pem -out database-service-server-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=database-service"
# Use CA's private key to sign database-service-server's CSR and get back the signed certificate
openssl x509 -req -in database-service-server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out database-service-server-cert.pem -extfile ./altNames_database-service.cnf
gpg -e -a -r dockerRunning database-service-server-key.pem
mv database-service-server-cert.pem ../database-service/certs/
mv database-service-server-key.pem.asc ../database-service/certs/

# strava-service
# Generate strava-service-server private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout strava-service-server-key.pem -out strava-service-server-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=strava-service"
# Use CA's private key to sign strava-service-server's CSR and get back the signed certificate
openssl x509 -req -in strava-service-server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out strava-service-server-cert.pem -extfile ./altNames_strava-service.cnf
gpg -e -a -r dockerRunning strava-service-server-key.pem
mv strava-service-server-cert.pem ../strava-service/certs/
mv strava-service-server-key.pem.asc ../strava-service/certs/

# runner
# Generate runner-server private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout runner-server-key.pem -out runner-server-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=runner"
# Use CA's private key to sign runner-server's CSR and get back the signed certificate
openssl x509 -req -in runner-server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out runner-server-cert.pem -extfile ./altNames_runner.cnf
gpg -e -a -r dockerRunning runner-server-key.pem
mv runner-server-cert.pem ../runner/certs/
mv runner-server-key.pem.asc ../runner/certs/

# postgres
# Generate postgres-server private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout postgres-key.pem -out postgres-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=postgres"
# Use CA's private key to sign postgres' CSR and get back the signed certificate
openssl x509 -req -in postgres-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out postgres-cert.pem -extfile ./altNames_postgres.cnf
gpg -e -a -r dockerRunning postgres-key.pem
cp postgres-cert.pem ../database-service/certs/
cp postgres-key.pem.asc ../database-service/certs/
mv postgres-cert.pem ../postgres/certs/
mv postgres-key.pem.asc ../postgres/certs/

# running_app
# Generate running_app private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout running_app-key.pem -out running_app-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=running_app"
# Use CA's private key to sign running_app's CSR and get back the signed certificate
openssl x509 -req -in running_app-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out running_app-cert.pem -extfile ./altNames_running_app.cnf
gpg -e -a -r dockerRunning running_app-key.pem
mv running_app-cert.pem ../running_app/certs/
mv running_app-key.pem.asc ../running_app/certs/

# http
# Generate http_gateway-server private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout http_gateway-server-key.pem -out http_gateway-server-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=http_gateway"
# Use CA's private key to sign http_gateway-server's CSR and get back the signed certificate
openssl x509 -req -in http_gateway-server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out http_gateway-server-cert.pem -extfile ./altNames_http_gateway.cnf
gpg -e -a -r dockerRunning http_gateway-server-key.pem
mv http_gateway-server-cert.pem ../httpGateway/certs/
mv http_gateway-server-key.pem.asc ../httpGateway/certs/