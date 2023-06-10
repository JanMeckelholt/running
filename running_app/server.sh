#!/bin/bash
PORT=443

echo 'Server starting on port' $PORT '...'
echo $(pwd)
cd /running_app/run
python3 server.py