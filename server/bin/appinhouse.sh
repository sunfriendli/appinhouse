#!/bin/bash
### BEGIN INIT INFO
# Provides:          appinhouse
# Required-Start:    $local_fs $remote_fs $network
# Required-Stop:     $local_fs $remote_fs $network
# Should-Start:      $named
# Should-Stop:       $named
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Start appinhouse
# Description:       Start appinhouse.
### END INIT INFO

usage()
{
    echo "Usage: ${0##*/} {status|start|stop|restart}"
    exit 1
}
#set -x

[ $# -gt 0 ] || usage

ACTION=$1
DEFAULT=/etc/default/appinhouse
NAME=appinhouse
APPINHOUSE_USER=appinhouse

if [ -f "$DEFAULT" ]; then
    . "$DEFAULT"
else
    exit 0
fi


PID_FILE="/var/run/$NAME.pid"

running()
{
    if [ -f "$1" ]
    then
        local PID=$(cat "$1" 2>/dev/null)
        RETVAL="$?"
        kill -0 "$PID" 2>/dev/null
        return $RETVAL
    fi
    rm -f "$1"
    RETVAL=1
    return $RETVAL
}

stop()
{
    start-stop-daemon --stop --quiet --oknodo --retry=TERM/30/KILL/5 --pidfile $PID_FILE
    RETVAL="$?"
    rm -f $PID_FILE
    return "$RETVAL"
}


start()
{
    start-stop-daemon --start --quiet --background --m --pidfile $PID_FILE --chuid $APPINHOUSE_USER --exec $ARTIFACT
    RETVAL="$?"
    return "$RETVAL"
}

##################################################
# Do the action
##################################################

case "$ACTION" in
    status)
        running "$PID_FILE"
        case "$?" in
            0) echo "$NAME running";;
            *) echo $NAME not runuing;; 
        esac
        ;;
    start)
	    start
        case "$?" in
            0) echo "Started $NAME";;
            1) echo "Already Running $(cat $PID_FILE)!";;  
            2) echo "Not Started  $NAME";;
        esac
        ;;
    stop)
        stop
        case "$?" in
            0) echo "$NAME stopped";;
            1) echo "$NAME already not running.";;
            2) echo "$NAME not stopped";;
        esac
        ;;
    restart)
        stop
        case "$?" in
            0|1)
                start
                case "$?" in
                    0) echo "Restarted $NAME" ;;
                    1) echo  Old process is still running $(cat $PID_FILE)!;;
                    *) echo "Restart Fail!";; # Failed to start
                esac
                ;;
            *)
        esac
        ;;
    *)
        usage
        ;;
esac

exit $RETVAL
