#!/bin/bash
path=${1:-"/home/stressTest/us/yure"}
platform=${2:-"linux"}

echo "start ..."

#machine1="10.65.8.43"
machine2="10.128.78.14"
machine3="10.128.78.16"
machine4="10.128.78.19"
#machine5="10.128.71.69"
#machine6="10.128.71.70"
#machine7="10.128.71.80"

echo "path=$path, platform=$platform"

if [ $platform == "mac" ]
then
  go build .
else
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main.out
fi

echo "start packaging and upload files ..."

tar czvf stpCase.tar.gz ./main.out ./data/* ./dummy.py

#scp stpCase.tar.gz "root@"${machine1}":"${path}
scp stpCase.tar.gz "root@"${machine2}":"${path}
scp stpCase.tar.gz "root@"${machine3}":"${path}
scp stpCase.tar.gz "root@"${machine4}":"${path}
#scp stpCase.tar.gz "root@"${machine5}":"${path}
#scp stpCase.tar.gz "root@"${machine6}":"${path}
#scp stpCase.tar.gz "root@"${machine7}":"${path}

echo "packaging and upload files end ..."