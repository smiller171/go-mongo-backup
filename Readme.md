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
  -e "MONGOHOST=$(/sbin/ip route|awk '/default/ { print  $3}')" \
  openwhere/go-mongo-backup

```
This sets the mongo address to the IP of the Docker host.

## API Reference

To start a backup:  
`POST` /v0/dump
```json
{
    "bucket": "mybucket",
    "path": "/path/to/backup/dir/"
}
```

## Tests

ToDo

## Contributions

Pull requests accepted.

## License

MIT
