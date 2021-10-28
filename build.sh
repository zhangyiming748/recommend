#!/usr/bin/env bash
export GOROOT=/usr/local/go
export GOBIN=/usr/local/go/bin
#export GOPATH=/work/app/go_work/
export GOPATH=/work/app/go_work/go_depend_packs
export PATH=$PATH:$GOBIN

#go get github.com/gorilla/mux
#go get github.com/unrolled/render
#go get github.com/urfave/negroni
#go get github.com/garyburd/redigo/redis
#go get github.com/widuu/goini
#go get github.com/rs/cors
#go get github.com/pmylund/go-cache
go build

# tar for aliyun
mkdir -p /tmp/recommend-url
cp ./conf.ini /tmp/recommend-url/
cp ./recommend-url /tmp/recommend-url/
cd /tmp
tar -zcvf /tmp/recommend-url.tar.gz ./recommend-url
rm -rf recommend-url ;
#
#scp recommend-url.tar.gz root@47.95.243.101:/work/app/upload;
