#!/usr/bin/env bash

# 需要提前把配置都放入 config 文件夹中
kubectl create secret generic sign-config --from-file=config/