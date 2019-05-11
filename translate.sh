#!/bin/sh

mkdir -p output
outline=$(docker-compose exec font-outline-tools python3 getFontOutline.py $1) && \
docker-compose exec gopher-translator gopher-translator "${outline}" $1 > output/$1.txt