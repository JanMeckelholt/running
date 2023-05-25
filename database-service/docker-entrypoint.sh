#!/bin/sh
# Abort on any error (including if wait-for-it fails). set -e
# Wait for strava-service

/database-service/wait-for-it.sh "postgres:5432" -t 300
sleep 2
 # Run the main container command.
exec "$@"