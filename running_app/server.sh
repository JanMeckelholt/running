#!/bin/bash
PORT=5000

cd build/web/

echo 'Server starting on port' $PORT '...'
python3 -m http.server $PORT