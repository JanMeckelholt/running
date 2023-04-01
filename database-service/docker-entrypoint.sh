#!/bin/sh
# Abort on any error (including if wait-for-it fails). set -e
# Wait for strava-service

 /go/src/running/database-service/wait-for-it.sh "postgres:5432" -t 300

 # Run the main container command.
exec "$@"