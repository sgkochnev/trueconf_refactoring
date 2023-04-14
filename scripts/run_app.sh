#!/bin/bash

export HTTP_PORT="8080"
export JSON_STORE_NAME="./store/users.json"


go run ./cmd/http/main.go
