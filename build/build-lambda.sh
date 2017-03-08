#!/bin/bash
#
# Script to build and upload a docker image for this project
#
# $1 - tag of the image, defaults to latest
# $2 - namespace (develop, master or release)
#

set -e

NAME=mongoBackup
VERSION=${1}
BUCKET=${2}

DOCKER_VERSION=`docker --version | cut -f3 | cut -d '.' -f2`
[ ${DOCKER_VERSION} -lt 12 ] && TAG_FLAG='-f' || TAG_FLAG=''

cd lambda/mongo-backup
zip -r ../../${NAME}-${VERSION}.zip .
cd ../../
aws s3 cp ${NAME}-${VERSION}.zip s3://${BUCKET}/lambda/${NAME}-${VERSION}.zip
