#!/bin/sh
echo "doctl registry login"
doctl registry login
echo "docker-compose up"
docker-compose up