#!/bin/sh
# Abort on any error (including if wait-for-it fails). set -e
# Wait for database-service

 /go/src/running/populate-db/wait-for-it.sh "database-service:666" -t 300

 # Run the main container command.
exec "$@"