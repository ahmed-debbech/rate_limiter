#!/bin/bash

#usage: bash stress.sh MAX_REQS

echo "max reqs for each is: $1" 

n=$1
client1() {
  for (( i=0 ; i<$n ; i++ )); do
    curl -s -o /dev/null -w "client 1.1.1.1 $i %{http_code}\n" http://localhost:1700 -H "X-Forwarded-For: 1.1.1.1"
    if [ $(curl -s -o /dev/null -w "%{http_code}" http://localhost:1700 -H "X-Forwarded-For: 1.1.1.1") == "200\n" ]; then
        echo "client 1.1.1.1 $i 200"
    fi
  done
}

client2() {
  for (( i=0 ; i<$n ; i++ )); do
    if [ $(curl -s -o /dev/null -w "%{http_code}" http://localhost:1700 -H "X-Forwarded-For: 1.1.1.2") == "200\n" ]; then
        echo "client 1.1.1.2 $i 200"
    fi
  done
}

client3() {
  for (( i=0 ; i<$n ; i++ )); do
    curl -s -o /dev/null -w "client 1.1.1.3 $i %{http_code}\n" http://localhost:1700 -H "X-Forwarded-For: 1.1.1.3"
    if [ $(curl -s -o /dev/null -w "%{http_code}" http://localhost:1700 -H "X-Forwarded-For: 1.1.1.3") == "200\n" ]; then
        echo "client 1.1.1.3 $i 200"
    fi
  done
}

client1 &
client2 &
client3 &

wait

echo "both clients completed"