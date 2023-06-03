#!/bin/bash
PORT=5000

echo 'Server starting on port' $PORT '...'
echo $(pwd)
cd running_app
python3 server.py