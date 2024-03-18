#!/bin/bash
export DEBUG=true
export CONF_PATH=../../config/conf.yaml
export LOG_DIR=../../log
export MEASURE=enable


if [ -z "$1" ]
  then
    echo "Укажите название tools"
    exit 1
fi

MODULE=$1
FUNC_NAME=""

if [ -n "$2" ]
  then
    FUNC_NAME="--run $2"
    echo $FUNC_NAME
fi

cd ../tools/$MODULE
go test -v $FUNC_NAME