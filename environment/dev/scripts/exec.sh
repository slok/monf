#!/bin/bash

CURRENT_PATH=$(pwd)/$(dirname "${BASH_SOURCE[0]}")
source $CURRENT_PATH/config.sh
status_cmd="docker-compose -f $CURRENT_PATH/../docker-compose.yml -p monf ps"

if [ "$($status_cmd | grep monf_monf_run | awk '{print $3}')" == "Up" ]; then
    # get name
    container_name=$($status_cmd | grep monf_monf_run | awk '{print $1}')
    # Run
    docker exec -it $container_name $@
else
    docker-compose -f $CURRENT_PATH/../docker-compose.yml -p monf run --rm --service-ports monf $@
fi
