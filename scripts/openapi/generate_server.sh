#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"

docker run -ti --rm -v $(pwd):/mnt oapi-codegen oapi-codegen -generate types -o "/mnt/$output_dir/openapi_types.gen.go" -package "$package" "/mnt/api/openapi/$service.yaml"

docker run -ti --rm -v $(pwd):/mnt oapi-codegen oapi-codegen -generate chi-server -o "/mnt/$output_dir/openapi_api.gen.go" -package "$package" "/mnt/api/openapi/$service.yaml"

docker run -ti --rm -v $(pwd):/mnt oapi-codegen oapi-codegen -generate types -o "/mnt/internal/common/client/$service/openapi_types.gen.go" -package "$service" "/mnt/api/openapi/$service.yaml"

docker run -ti --rm -v /$(pwd):/mnt oapi-codegen oapi-codegen -generate client -o "/mnt/internal/common/client/$service/openapi_client_gen.go" -package "$service" "/mnt/api/openapi/$service.yaml"
