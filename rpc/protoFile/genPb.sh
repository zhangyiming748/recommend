#!/usr/bin/env bash
protoc --go_out=plugins=grpc:. article.proto
#protoc --go_out=plugins=grpc:. rank.proto
#protoc --go_out=plugins=grpc:. usercf.proto

#python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. es_test.proto

