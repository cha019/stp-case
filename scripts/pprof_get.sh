#!/bin/bash
is_num(){
    [[ "$1" =~ ^[1-9]+$ ]] && echo "check number ok" ||  echo "check number error"
}

process=$1

if [[ -z $process ]];
then
    echo "process is empty, pls input process"
    exit 1
fi

process_pid=$(ps -ef | grep -w $process | grep -v grep | grep -v bin | awk '{print $2}')
pprof_df=$(cat /proc/$process_pid/environ | tr '\0' '\n' | grep -i port | grep pprof | awk -F= '{print $2}')

wget -O trace.out localhost:$pprof_df/debug/pprof/trace
wget -O allocs.out localhost:$pprof_df/debug/pprof/allocs
wget -O profile.out localhost:$pprof_df/debug/pprof/profile
wget -O blocks.out localhost:$pprof_df/debug/pprof/blocks
wget -O cmdline.out localhost:$pprof_df/debug/pprof/cmdline
wget -O goroutine.out localhost:$pprof_df/debug/pprof/goroutine
wget -O heap.out localhost:$pprof_df/debug/pprof/heap
wget -O mutex.out localhost:$pprof_df/debug/pprof/mutex
wget -O threadcreate.out localhost:$pprof_df/debug/pprof/threadcreate

#execute cmd
#./pprof_get.sh ./info
#./pprof_get.sh ./account
#./pprof_get.sh ./payment_auth

#go tool pprof -http=:8090 profile.out
#smc -e test services download xx-xx-xx-xx -s /tmp/test/trace.out -d ./