#!/bin/bash

docker login -u $DOCKER_USER -p $DOCKER_PASS
docker build -t juniorrodes/mtg .

docker push juniorrodes/mtg