#!/bin/bash

PROCESS=Docker

# enter your postwoman folder path
cd D:/developer/postwoman


# start up the docker engine 

# for windows, simply put your path to docker desktop, the default is the following:
"C:\Program Files\Docker\Docker\Docker Desktop.exe"

# for mac
# open -a Docker


# wait for docker to start

# for windows
while [ ! "$(docker info | grep -q error)" ]
do
    sleep 1
done

echo "Docker"

# for mac
# while [ ! "$(ps aux | grep -v grep | grep -c $PROCESS)" -gt 0 ]
# do
#     sleep 1
# done


# start up postwoman

# docker-compose up
