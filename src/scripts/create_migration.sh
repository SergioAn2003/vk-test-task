#!/bin/bash

if [ -z "$1" ]
  then
    echo "Укажите название миграции"
    exit 1
fi

NAME=$1

migrate create -ext sql -dir ../docker/migrate -seq $NAME