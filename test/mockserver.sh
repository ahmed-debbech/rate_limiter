#!/bin/bash

PORT=8083

while true; do
  echo "Listening on port $PORT..."
    echo -e "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 2\r\n\r\nOK" | nc -l $PORT;

done
