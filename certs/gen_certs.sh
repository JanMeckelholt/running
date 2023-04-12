#!/usr/bin/env bash

rm *.pem
rm ../database-service/certs/*.pem
rm ../strava-service/certs/*.pem
rm ../postgres/certs/*.pem
rm ../runner/certs/*.pem
rm ../populate-db/certs/*.pem
rm ../httpGateway/certs/*.pem
rm ../running_app/certs/*.pem

# Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=DE/ST=x/L=x/O=x/CN=www.janmeckelholt.de" 
cp ca-cert.pem ../database-service/certs/
cp ca-cert.pem ../strava-service/certs/
cp ca-cert.pem ../postgres/certs/root.crt
cp ca-cert.pem ../runner/certs/
cp ca-cert.pem ../populate-db/certs/
cp ca-cert.pem ../httpGateway/certs/
cp ca-cert.pem ../running_app/certs/

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

# runner
# Generate runner-server private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout runner-server-key.pem -out runner-server-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=runner"
# Use CA's private key to sign runner-server's CSR and get back the signed certificate
openssl x509 -req -in runner-server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out runner-server-cert.pem -extfile ./altNames_runner.cnf
mv runner-server-cert.pem ../runner/certs/
mv runner-server-key.pem ../runner/certs/

# postgres
# Generate postgres-server private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout postgres-key.pem -out postgres-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=postgres"
# Use CA's private key to sign postgres' CSR and get back the signed certificate
openssl x509 -req -in postgres-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out postgres-cert.pem -extfile ./altNames_postgres.cnf
cp postgres-cert.pem ../database-service/certs/postgres-cert.pem
cp postgres-key.pem ../database-service/certs/postgres-key.pem
chown 70:70 postgres-key.pem
chmod 600 postgres-key.pem
mv postgres-cert.pem ../postgres/certs/postgres-cert.pem
mv postgres-key.pem ../postgres/certs/postgres-key.pem 

# running_app
# Generate running_app private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout running_app-key.pem -out running_app-req.pem -subj "/C=DE/ST=x/L=x/O=x/CN=running_app"
# Use CA's private key to sign running_app's CSR and get back the signed certificate
openssl x509 -req -in running_app-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out running_app-cert.pem -extfile ./altNames_running_app.cnf
mv running_app-cert.pem ../running_app/certs/
mv running_app-key.pem ../running_app/certs/