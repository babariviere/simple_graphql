#!/bin/sh

USAGE="usage: run.sh <service>

Services:
    graphql     graphql backend
    projects    projects microservice
    users       users microservice"

case $1 in
    graphql)
        echo "running graphql server";;
    projects)
        echo "running projects microservice";;
    users)
        echo "running users microservice";;
    *)
        echo "$USAGE";;
esac
