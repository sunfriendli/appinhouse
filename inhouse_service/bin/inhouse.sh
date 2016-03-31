#!/bin/bash

usage()
{
    echo "Usage: ${0##*/} {status|start|stop|restart}"
    exit 1
}

[ $# -gt 0 ] || usage

ACTION=$1
ARTIFACT="../inhouse_service"

PID_FILE="$ARTIFACT.pid"

running()
{
    if [ -f "$1" ]
    then
        local PID=$(cat "$1" 2>/dev/null) || return 1
        kill -0 "$PID" 2>/dev/null
        return
    fi
    rm -f "$1"
    return 1
}


stop()
{
    if [ ! -f $PID_FILE ];
    then
        echo "File $PID_FILE not found."
        return
    fi
    
    local PID=$(cat $PID_FILE)
    echo -n "Stopping $ARTIFACT pid=$PID "
    kill -TERM $PID
    local TIMEOUT=30
    while running "$PID_FILE"; do
        if (( TIMEOUT-- == 0 )); then
          echo -n "TIMEOUT - stopping with KILL signal "
          kill -KILL $PID
        fi
        sleep 1
        echo -n ". "
    done
    echo ''
    echo "Stopped $ARTIFACT"
}


start()
{
    nohup  ./$ARTIFACT > /dev/null 2>&1 &
    local PID=$!
    disown $PID
    echo $PID > "$PID_FILE"
    echo "Started $ARTIFACT pid=$PID"
}

##################################################
# Do the action
##################################################

case "$ACTION" in
    status)
        if running "$PID_FILE"
        then
            echo "$ARTIFACT running pid=$(< "$PID_FILE")"
            exit 0
        fi
        exit 1
        ;;
    start)
        if running $PID_FILE
        then
            echo "Already Running $(cat $PID_FILE)!"
            exit 1
        fi
        echo "Starting $ARTIFACT"
        start
        ;;
    stop)
        if running $PID_FILE
        then
            echo "Stopping $ARTIFACT"
            stop
            exit 0
        fi
        echo "$ARTIFACT not running."
        exit 1
        ;;
    restart)
        echo "Restarting $ARTIFACT"
        stop
        start
        ;;
    *)
        usage
        ;;
esac
