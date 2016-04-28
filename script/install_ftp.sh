#!/bin/bash

usage()
{
    echo "Usage: ${0##*/} ftp_root_dir passwd_postfix"
	echo "example :${0} /var/vsftpd pwpostfix"
	echo "usernames see appinhouse/server/conf/app.conf app_names"
	echo "create ftp user and passwd:"
	echo "name1:name1pwpostfix"
	echo "name2:name2pwpostfix"
    exit 1
}
FTP_ROOT_DIR=$1
PASSWD_POSTDIX=$2

if [ -z "$PASSWD_POSTDIX" ]; then
    echo "password  postfix is null"
    usage
fi
if [ -z "$FTP_ROOT_DIR" ]; then
    echo "ftp root dir is null"
    usage
fi
echo "install vsftpd..."
sudo apt-get -y install vsftpd

#create user
echo "create user..."
INHOUSE_DEFAULT_LOCATION="/home/appinhouse"
GIT_DIR=appinhouse_git
APPINHOUSE_HOME=$INHOUSE_DEFAULT_LOCATION/$GIT_DIR/src/appinhouse/server
APPS=$(echo "$(grep 'app_names' $APPINHOUSE_HOME/conf/app.conf)" |sed 's/ //g'|cut -c 11-)
ANDROID_DEV_PATH=dev/android/data
IOS_DEV_PATH=dev/ios/data
ANDROID_RELEASE_PATH=release/android/data
IOS_RELEASE_PATH=release/ios/data
arr=(${APPS//;/ })

for i in ${arr[@]}
do
    if [ ! -d $FTP_ROOT_DIR$i ]; then
        sudo mkdir -p $FTP_ROOT_DIR/$i/$ANDROID_DEV_PATH
        sudo mkdir -p $FTP_ROOT_DIR/$i/$IOS_DEV_PATH
        sudo mkdir -p $FTP_ROOT_DIR/$i/$ANDROID_RELEASE_PATH
        sudo mkdir -p $FTP_ROOT_DIR/$i/$IOS_RELEASE_PATH
        sudo useradd $i -d $FTP_ROOT_DIR/$i -s /bin/false
        echo "sudopsw" | sudo -S echo $i:$i$PASSWD_POSTDIX | sudo chpasswd
        sudo chown -R $i:$i  $FTP_ROOT_DIR/$i
    fi
done
#vsftpd conf
echo "copy conf..."
FTP_CONF=$INHOUSE_DEFAULT_LOCATION/$GIT_DIR/src/appinhouse/conf/ftp
sudo mkdir -p /etc/vsftpd
sudo cp  -R $FTP_CONF/vsftpd_user_conf /etc/vsftpd/
sudo cp   $FTP_CONF/vsftpd.conf /etc
sudo cp $FTP_CONF/vsftpd /etc/pam.d/
sudo service vsftpd restart