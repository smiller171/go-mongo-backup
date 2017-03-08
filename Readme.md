# go-mongo-backup
---

## Description

This application runs alongside Mongo to create a simple API endpoint that can be called for backup operations.

## Example

```bash
curl -XPOST 'http://localhost:8080/v0/dump' -d \
'{
    "bucket": "org-mongodb-snapshots",
    "path": "/dev/"
}'
```

This creates a timestamped file in your S3 bucket at the path you specify. This file is compatible with `mongorestore`, using the `--archive` option

## Motivation

I wanted to be able to use a Lambda scheduled job to back up Mongo to S3, but using mongodump in python required more than 1500MB of memory on my dataset and was very slow.

## Installation

If mongo is running in another container:

```bash
docker pull openwhere/go-mongo-backup
docker run --name mongo-backup --link mongo -d -p 8080:8080 openwhere/go-mongo-backup
```

If mongo is running as a native package on the host:

```bash
docker pull openwhere/go-mongo-backup
docker run --name mongo-backup -d -p 8080:8080 \
  -e "MONGOHOST=$(ifconfig | grep -E '([0-9]{1,3}\.){3}[0-9]{1,3}' | grep -v 127.0.0.1 | awk '{ print $2 }' | cut -f2 -d: | head -n1)" \
  openwhere/go-mongo-backup

```
This sets the mongo address to the IP of the Docker host.

## Configuration
This Lambda is configured with the following environment variables:

    URL="http://mongodb.mydomain.com:8080"
    BUCKET="mybucket"
    ROOT_PATH="/backup/path/"

## API Reference

To start a backup:  
`POST` /v0/dump
```json
{
    "bucket": "mybucket",
    "path": "/path/to/backup/dir/",
    "region": "us-east-1"
}
```

## Lambda build
1.  Test the lambda function
  1.  Set environment variables
            export URL="http://mongodb.mydomain.com:8080"
            export BUCKET="mybucket"
            export ROOT_PATH="/path/for/backup/storage/"
            export AWS_DEFAULT_REGION="us-east-1"
  2.  Run lambda
            python lambda/mongo-backup/mongo-snapshot.py
2.  create zip
        ./build/lambda-build.sh ${VERSION} ${BUCKET}
3.  Create Lambda function using zip file. **Must pass environment variables** (Automated in pod config)

## Container build
1. $(aws ecr get-login)
2. ./build/build-docker.sh ${VERSION} ${NAMESPACE}

## Tests

ToDo

## Contributions

Pull requests accepted.

## License

MIT
