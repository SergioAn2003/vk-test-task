#!/bin/bash
export DEBUG=true
export CONF_PATH=../../../../config/conf.yaml
export CUSTOM_TEMPLATE_DIR=../../../../custom_templates


if [ -z "$1" ]
  then
    echo "Укажите название usecase"
    exit 1
fi

MODULE=$1
FUNC_NAME=""

if [ -n "$2" ]
  then
    FUNC_NAME="--run $2"
    echo $FUNC_NAME
fi

cd ../internal/usecase/test/$MODULE
go test -v $FUNC_NAME