#!/bin/bash

#
# Install inhouse_sevice on Ubuntu Server, and configure it to be a daemon service.
#


INHOUSE_PASSWD=$1
if [ -z "$INHOUSE_PASSWD" ]; then
    echo "password is null"
    exit 1
fi
CUR_USER=$LOGNAME
USER=appinhouse
INHOUSE_DEFAULT_LOCATION="/home/appinhouse"
BEE_DIR=bee
GIT_DIR=appinhouse_git
TARGET=appinhouse_server
APPNAME=inhouse_service
APPINHOUSE_WEB=appinhouse_web

#create inhouse home if not exists 

if [ ! -d $INHOUSE_DEFAULT_LOCATION ]; then
  sudo mkdir $INHOUSE_DEFAULT_LOCATION
fi
id $USER >& /dev/null
if [ $? -ne 0 ]
then
   echo 'add user $USER...'
   sudo useradd $USER -s /bin/bash
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
APPINHOUSE_HOME=$INHOUSE_DEFAULT_LOCATION/$GIT_DIR/src/appinhouse/inhouse_service
cd $APPINHOUSE_HOME
echo 'get and build'
go get -v

go build

echo 'package...'
bee pack -o "$INHOUSE_DEFAULT_LOCATION" -exr pack.sh

echo 'deploy...'
cd $INHOUSE_DEFAULT_LOCATION
if [ ! -d $TARGET ];
    then
        echo "File $TARGET not found."
        mkdir $TARGET
fi
tar -zxvf $APPNAME.tar.gz -C $TARGET
echo "deploy file in $INHOUSE_DEFAULT_LOCATION/$TARGET" 
mkdir $APPINHOUSE_WEB
cp -R $INHOUSE_DEFAULT_LOCATION/$GIT_DIR/src/appinhouse/web/static/* $APPINHOUSE_WEB/

#make appinhouse to a service
FTP_ROOT_DIR=$(echo "$(grep 'ftp_root_dir' $APPINHOUSE_HOME/conf/app.conf)" |sed 's/ //g'|cut -c 14-)
APPS=$(echo "$(grep 'app_names' $APPINHOUSE_HOME/conf/app.conf)" |sed 's/ //g'|cut -c 11-)
ANDROID_DEV_PATH=dev/android/data
IOS_DEV_PATH=dev/ios/data
ANDROID_RELEASE_PATH=release/android/data
IOS_RELEASE_PATH=release/ios/data
arr=(${APPS//;/ })

for i in ${arr[@]}
do
    if [ ! -d $FTP_ROOT_DIR$i ]; then
        sudo mkdir -p $FTP_ROOT_DIR$i/$ANDROID_DEV_PATH
        sudo mkdir -p $FTP_ROOT_DIR$i/$IOS_DEV_PATH
        sudo mkdir -p $FTP_ROOT_DIR$i/$ANDROID_RELEASE_PATH
        sudo mkdir -p $FTP_ROOT_DIR$i/$IOS_RELEASE_PATH
    fi
done

sudo cp $APPINHOUSE_HOME/bin/appinhouse.sh /etc/init.d/appinhouse
sudo chmod +x /etc/init.d/appinhouse
echo "ARTIFACT=$INHOUSE_DEFAULT_LOCATION/$TARGET/inhouse_service"       | sudo tee -a /etc/default/appinhouse
echo "APPINHOUSE_USER=$USER"                      | sudo tee -a /etc/default/appinhouse
sudo chown -R $USER:$USER $INHOUSE_DEFAULT_LOCATION
sudo update-rc.d appinhouse defaults
sudo service appinhouse start
