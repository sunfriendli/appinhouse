#!/usr/bin/env bash
echo 'start'

cd ../../../..

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
HOME=env | grep ^HOME= | cut -c 6-
TARGET=appinhouse_server

export GOPATH="$CURDIR"

cd src/appinhouse/service

echo 'get and build'
go get -v

go build 

echo 'package...'
bee pack -o "$HOME"

export GOPATH="$OLDGOPATH"

echo 'deploy...'
cd ~
if [ ! -d $TARGET ];
    then
        echo "File $TARGET not found."
		mkdir $TARGET
fi
tar -zxvf service.tar.gz -C $TARGET

echo 'finished'
