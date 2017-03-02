import logging
import json
import requests
import sys
from os import environ
from os import path

logger = logging.getLogger()
logger.setLevel(logging.INFO)
# logger.setLevel(logging.DEBUG)


# This is the method that will be registered
# with Lambda and run on a schedule
def handler(event={}, context={}):
    logger.info("started")

    domain = environ['URL'].strip('/ ')
    bucket = environ['BUCKET'].strip()
    basePath = environ['ROOT_PATH'].strip()
    region = environ['AWS_DEFAULT_REGION']

    logger.info("bucket is " + bucket)
    logger.info("base_path is " + basePath)

    logger.info("setting mongodump json")
    target = {
        "bucket": bucket,
        "path": path.join("/", basePath, ""),
        "region": region
    }
    logger.info("mongodump json is " + json.dumps(target))

    logger.info("setting url path")
    url = path.join(domain, "v0/dump")
    logger.info("url path is " + url)

    # trigger dump
    logger.info("triggering dump")
    try:
        response = requests.post(
            url,
            data=json.dumps(target)
            )
        response.raise_for_status()
    except requests.exceptions.RequestException as e:
        logger.error(e)
        logger.error(response.content)
        sys.exit(1)

    logger.info(response.content)
    logger.info("new snapshot started at " + url)


# If being called locally, just call handler
if __name__ == '__main__':
    import os
    import json
    import sys

    logging.basicConfig()
    event = {}

    # TODO if argv[1], read contents, parse into json
    if len(sys.argv) > 1:
        input_file = sys.argv[1]
        with open(input_file, 'r') as f:
            data = f.read()
        event = json.loads(data)

    result = handler(event)
    output = json.dumps(
        result,
        sort_keys=True,
        indent=4,
        separators=(',', ':')
    )
    logger.info(output)
