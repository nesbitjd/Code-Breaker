#!/bin/bash
sleep 5
curl --request POST \
  --url http://localhost:8080/api/v1/word \
  --header 'Content-Type: application/json' \
  --data '{
	"word": "frugalicious"
}'