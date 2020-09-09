#!/usr/bin/env bash


rm -rf build

# echo-service
mkdir -p build/echo-service/conf
cp -f echo-service/cmd/echo-server/conf/config.yaml build/echo-service/conf/config.yaml
go build -o build/echo-service/echo-server  echo-service/cmd/echo-server/main.go


echo "build ok"
#
#CROSS_HOST=114.67.106.133
#CROSS_PATH=/root/projects
#scp -r build/* root@$CROSS_HOST:$CROSS_PATH
