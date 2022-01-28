#!/usr/bin/env bash
# 该脚本用于生成 *.pb.go

modules=$(ls -l | awk '/^d/ {print $NF}')
for module in $modules
do
  protoc --proto_path=.:/usr/local/include \
         --go_out=. --go_opt=paths=source_relative \
         --micro_out=. --micro_opt=paths=source_relative \
         ./$module/*.proto
done
