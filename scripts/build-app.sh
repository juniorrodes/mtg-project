#!/bin/bash
set -e

TAG = $(git describe --tags &(git rev-list --tags --max-count=1))
TAG = $(echo $TAG | sed 's/^v//') >> "$GITHUB_ENV"

docker login -u $DOCKER_USER -p $DOCKER_PASS
docker build -t juniorrodes/mtg:$TAG .

docker push juniorrodes/mtg:$TAG