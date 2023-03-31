#!/bin/bash
set -e

readonly service="$1"

docker run --rm --env "JAVA_OPTS=-Dlog.level=error" -v "${PWD}:/local" \
  "openapitools/openapi-generator-cli:v4.3.0" generate --skip-validate-spec \
  -i "/local/api/openapi/$service.yaml" \
  -g javascript \
  -o "/local/web/$service"