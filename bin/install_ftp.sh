#!/bin/bash

PASSWD_POSTDIX=$1
if [ -z "$PASSWD_POSTDIX" ]; then
    echo "password  postfix is null"
    exit 1
fi

INHOUSE_DEFAULT_LOCATION="/home/appinhouse"
GIT_DIR=appinhouse_git

sudo apt-get -y install vsftpd

#create user
APPINHOUSE_HOME=$INHOUSE_DEFAULT_LOCATION/$GIT_DIR/src/appinhouse/inhouse_service
FTP_ROOT_DIR=$(echo "$(grep 'ftp_root_dir' $APPINHOUSE_HOME/conf/app.conf)" |sed 's/ //g'|cut -c 14-)
APPS=$(echo "$(grep 'app_names' $APPINHOUSE_HOME/conf/app.conf)" |sed 's/ //g'|cut -c 11-)
ANDROID_DEV_PATH=dev/android/data
IOS_DEV_PATH=dev/ios/data
ANDROID_RELEASE_PATH=release/android/data
IOS_RELEASE_PATH=release/ios/data
arr=(${APPS//;/ })

for i in ${arr[@]}
do
    if [ -d $FTP_ROOT_DIR$i ]; then
        sudo useradd $i -d $FTP_ROOT_DIR$i -s /bin/false
        echo "sudopsw" | sudo -S echo $i:$i$PASSWD_POSTDIX | sudo chpasswd
        sudo chown -R $i:$i  $FTP_ROOT_DIR$i
    else
        sudo mkdir -p $FTP_ROOT_DIR$i/$ANDROID_DEV_PATH
        sudo mkdir -p $FTP_ROOT_DIR$i/$IOS_DEV_PATH
        sudo mkdir -p $FTP_ROOT_DIR$i/$ANDROID_RELEASE_PATH
        sudo mkdir -p $FTP_ROOT_DIR$i/$IOS_RELEASE_PATH
        sudo useradd $i -d $FTP_ROOT_DIR$i -s /bin/false
        echo "sudopsw" | sudo -S echo $i:$i$PASSWD_POSTDIX | sudo chpasswd
        sudo chown -R $i:$i  $FTP_ROOT_DIR$i
    fi
done
#vsftpd conf
FTP_CONF=$INHOUSE_DEFAULT_LOCATION/$GIT_DIR/src/appinhouse/ftp
sudo mkdir -p /etc/vsftpd
sudo cp  -R $FTP_CONF/vsftpd_user_conf /etc/vsftpd/
sudo cp   $FTP_CONF/vsftpd.conf /etc
sudo cp $FTP_CONF/vsftpd /etc/pam.d/
sudo service vsftpd restart