#!/usr/bin/env bash

export ROOT=..
source build_only.sh

echo 'RUN'
$ROOT/bin/$BIN $@
