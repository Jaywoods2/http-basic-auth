#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

mv main httpbasicauth

docker build -t registry.cn-shenzhen.aliyuncs.com/jaywoods/httpbasicauth .

docker push  registry.cn-shenzhen.aliyuncs.com/jaywoods/httpbasicauth:latest