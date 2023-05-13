#!/bin/bash
PORT=5000

mv server.py ./build/web/
cd build/web/


echo 'Server starting on port' $PORT '...'
python3 server.py