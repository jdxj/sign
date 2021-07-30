#!/usr/bin/env bash

pid=$(ps -ef | grep "apiserver" | grep -v "grep" | awk '{print $2}')
if [ -n "${pid}" ]; then
  kill "$pid"
fi

wd=/root/app/sign
jwd=/root/.jenkins/jobs/test/workspace

rm -f ${wd}/apiserver.out
cp ${jwd}/_output/build/apiserver.out ${wd}

nohup ${wd}/apiserver.out -f ${wd}/config.yaml >> ${wd}/nohup_apiserver.log 2>&1 &