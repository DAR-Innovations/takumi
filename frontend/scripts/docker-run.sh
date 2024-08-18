#!/bin/bash
DOCKER_BUILDKIT=1  docker build -t dar-innovations/takumi-frontend . 
docker run -d -p 4173:80 --rm --name takumi-frontend dar-innovations/takumi-frontend