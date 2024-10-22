#!/bin/bash

docker-compose -f ./mysql-init.yaml down
rm -r /var/mysql-docker/data
docker-compose -f ./mysql-init.yaml up -d