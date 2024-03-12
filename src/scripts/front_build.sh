#!/usr/bin/env bash

export ROOT=..

cd $ROOT
rm _nuxt -r
cd ../fronted
yarn build
cp ./dist/_nuxt -r ../appv3
cp ./dist/index.html ../appv3/templates/main.html

#cd ../app
#./start.sh