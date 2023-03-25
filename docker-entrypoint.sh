#!/bin/sh
# Abort on any error (including if wait-for-it fails). set -e
# Wait for strava-service

 /go/src/app/wait-for-it.sh "strava-servcie:666" -t 20

 # Run the main container command.
exec "$@"