#!/bin/bash

docker stop minimum-ca-db
docker rm minimum-ca-db

docker run -d \
    --name minimum-ca-db \
    -p 3306:3306 \
    -e TZ=Asia/Tokyo \
    -e MYSQL_ROOT_PASSWORD=test \
    -e MYSQL_USER=test \
    -e MYSQL_PASSWORD=test \
    minimum-ca-db:latest
