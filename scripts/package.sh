#!/bin/bash
path=${1:-"/home/stressTest/xx"}
platform=${2:-"linux"}

echo "start ..."

machine1="xx.xx.xx.xx"

echo "path=$path, platform=$platform"

if [ $platform == "mac" ]
then
  go build .
else
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main.out
fi

echo "start packaging and upload files ..."

tar czvf stpCase.tar.gz ./main.out ./data/* ./dummy.py

scp stpCase.tar.gz "root@"${machine1}":"${path}

echo "packaging and upload files end ..."