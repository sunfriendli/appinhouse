#!/bin/bash

#
# Install inhouse_sevice on Ubuntu Server, and configure it to be a daemon service.
#

usage()
{
    echo "Usage: ${0##*/} password redis_address redis_password(optional)"
    exit 1
}

INHOUSE_PASSWD=$1
REDIS_ADDR=$2
REDIS_PASSWORD=$3
USER=appinhouse

if [ -z "$INHOUSE_PASSWD" ]; then
    echo "create a user by defaul.username is $USER"
   usage
fi
if [ -z "$REDIS_ADDR" ]; then
   usage
fi

CUR_USER=$LOGNAME

INHOUSE_DEFAULT_LOCATION="/home/appinhouse"
BEE_DIR=bee
GIT_DIR=appinhouse_git
TARGET=appinhouse_server
APPNAME=server
APPINHOUSE_WEB=appinhouse_web

id $USER >& /dev/null
if [ $? -ne 0 ]
then
   echo 'add user $USER...'
   sudo useradd $USER -m -s /bin/bash
   echo "sudopsw" | sudo -S echo $USER:$INHOUSE_PASSWD | sudo chpasswd
   sudo chown -R $CUR_USER:$CUR_USER  $INHOUSE_DEFAULT_LOCATION
fi

#install go
echo 'install go...'
sudo add-apt-repository ppa:ubuntu-lxc/lxd-stable
sudo apt-get update
sudo apt-get -y install golang

#install git
echo 'install git...'
sudo apt-get -y install git

#install bee
echo 'install bee...'
mkdir $INHOUSE_DEFAULT_LOCATION/$BEE_DIR
export GOPATH=$INHOUSE_DEFAULT_LOCATION/$BEE_DIR
cd $GOPATH
go get github.com/beego/bee
cd bin
sudo  cp bee /usr/bin

 
# downlown source
echo 'downlown source...'
mkdir  $INHOUSE_DEFAULT_LOCATION/$GIT_DIR
cd $INHOUSE_DEFAULT_LOCATION/$GIT_DIR
mkdir pkg
mkdir bin
mkdir src
cd src
git clone https://github.com/rog2/appinhouse.git



# pack appinhouse
export GOPATH=$INHOUSE_DEFAULT_LOCATION/$GIT_DIR
APPINHOUSE_HOME=$INHOUSE_DEFAULT_LOCATION/$GIT_DIR/src/appinhouse/server
cd $APPINHOUSE_HOME
echo 'get and build'
go get -v

go build -o appinhouse

echo 'package...'
bee pack -o "$INHOUSE_DEFAULT_LOCATION" -exr pack.sh -exr server

echo 'deploy...'
cd $INHOUSE_DEFAULT_LOCATION
if [ ! -d $TARGET ];
    then
        echo "mkdir $TARGET"
        mkdir $TARGET
fi
tar -zxvf $APPNAME.tar.gz -C $TARGET
echo "deploy file in $INHOUSE_DEFAULT_LOCATION/$TARGET" 
mkdir $APPINHOUSE_WEB
cp -R $INHOUSE_DEFAULT_LOCATION/$GIT_DIR/src/appinhouse/web/static/* $APPINHOUSE_WEB

echo 'set env...'
REDIS_CONF=$(echo "$(grep 'conf_dir' $APPINHOUSE_HOME/conf/app.conf)" |sed 's/ //g'|cut -c 10-)
REDIS_ADDR_NAME=$(echo "$(grep 'env_addr_name' $APPINHOUSE_HOME/conf/app.conf)" |sed 's/ //g'|cut -c 15-)
REDIS_PASSWORD_NAME=$(echo "$(grep 'env_password_name' $APPINHOUSE_HOME/conf/app.conf)" |sed 's/ //g'|cut -c 19-)
REDIS_DIR=$(dirname $REDIS_CONF)
if [ ! -d $REDIS_DIR ];
    then
        echo "mkdir $REDIS_DIR"
        sudo mkdir -p $REDIS_DIR
fi
sudo echo ""  | sudo tee $REDIS_CONF
echo "$REDIS_ADDR_NAME=$REDIS_ADDR"       | sudo tee -a $REDIS_CONF
if [ -n "$REDIS_PASSWORD" ]; then
  echo "$REDIS_PASSWORD_NAME=$REDIS_PASSWORD"    | sudo tee -a $REDIS_CONF
fi

#make appinhouse to a service
sudo cp $APPINHOUSE_HOME/bin/appinhouse.sh /etc/init.d/appinhouse
sudo chmod +x /etc/init.d/appinhouse
echo "ARTIFACT=$INHOUSE_DEFAULT_LOCATION/$TARGET/appinhouse"       | sudo tee -a /etc/default/appinhouse
echo "APPINHOUSE_USER=$USER"                      | sudo tee -a /etc/default/appinhouse
sudo chown -R $USER:$USER $INHOUSE_DEFAULT_LOCATION
sudo update-rc.d appinhouse defaults
sudo service appinhouse start
