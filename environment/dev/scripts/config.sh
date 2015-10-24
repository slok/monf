#!/bin/bash

IMAGE_NAME="slok/monf"
CONTAINER_NAME="monf_run"
export HOSTIP=$(ip addr | grep 'state UP' -A2 | tail -n1 | awk '{print $2}' | cut -f1  -d'/')