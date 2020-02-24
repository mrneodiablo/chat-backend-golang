#!/bin/bash

GOPATH="/dongvt_project_exchange_pet/server/jenkins/workspace/chat-server"
export GOPATH=${GOPATH}
export PATH=$PATH:${GOPATH}/bin



SOURCE="/dongvt_project_exchange_pet/server/jenkins/workspace/build-gateway-chat"
DESTI=${GOPATH}/src/hope-pet-chat-backend

BINARYNAME="pet-chatting"
PID_PATH_NAME=${GOPATH}/run/server.pid


case $1 in
    start)

        # install go govendor
        rm -rf ${DESTI}/*
        cp -R ${SOURCE}/*  ${DESTI}/
        cd ${DESTI}
        go get -u github.com/kardianos/govendor


        # clean
        go clean
        rm -rf ${DESTI}/${BINARYNAME}

        # build
        echo "build golang ..."
        govendor sync -v
        go build -o ${BINARYNAME} -v main.go

        echo "Starting SERVER CHAT ..."
        if [ ! -f $PID_PATH_NAME ]; then
            chmod +x ${BINARYNAME}
            nohup  ./${BINARYNAME} > ${GOPATH}/logs/server-chat.`date +"%Y%m%d%H%M"`.log &

            echo $! > $PID_PATH_NAME
            echo "SERVER CHAT started ..."
        else
            echo "SERVER CHAT is already running ..."
        fi
    ;;
    stop)
        if [ -f $PID_PATH_NAME ]; then
            PID=$(cat $PID_PATH_NAME);
            echo "SERVER CHAT stoping ..."
            kill $PID;
            echo "SERVER CHAT stopped ..."
            rm $PID_PATH_NAME
        else
            echo "SERVER CHAT is not running ..."
        fi
    ;;
esac