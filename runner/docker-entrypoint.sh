#!/bin/sh
# Abort on any error (including if wait-for-it fails). set -e
# Wait for strava-service

 /go/src/running/runner/wait-for-it.sh "strava-service:666" -t 300

 # Run the main container command.
exec "$@"