#!/bin/bash

CURRENT_PATH=$(dirname "${BASH_SOURCE[0]}")
source $CURRENT_PATH/config.sh

docker-compose -f $CURRENT_PATH/../docker-compose.yml -p monf stop
