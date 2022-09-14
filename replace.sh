#!/bin/bash
path=${1:-"/home/stressTest/us/yure"}
platform=${2:-"linux"}

echo "start ..."

machine1="10.65.8.43"
#machine2="10.128.78.14"
#machine3="10.128.78.16"
#machine4="10.128.78.19"
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

scp main.out "root@"${machine1}":"${path}
#scp main.out "root@"${machine2}":"${path}
#scp main.out "root@"${machine3}":"${path}
#scp main.out "root@"${machine4}":"${path}
#scp main.out "root@"${machine5}":"${path}
#scp main.out "root@"${machine6}":"${path}
#scp main.out "root@"${machine7}":"${path}

echo "packaging and upload files end ..."
