#!/bin/bash

#
# Install inhouse_sevice on Ubuntu Server, and configure it to be a daemon service.
#

usage()
{
    echo "Usage: ${0##*/} password(optional)"
    exit 1
}

REDIS_PASSWORD=$1

#install redis
echo 'install redis...'
sudo apt-get -y install redis-server
if [ -n "$REDIS_PASSWORD" ]; then
   echo "requirepass $REDIS_PASSWORD" |sudo tee -a /etc/redis/redis.conf
fi
sudo service redis-server restart
