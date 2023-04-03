include .env

.PHONY: bootstrap
bootstrap: openapi

.PHONY: openapi
openapi: openapi_http openapi_js merge_openapi_docs sqlc

.PHONY: merge_openapi_docs
merge_openapi_docs:
	docker run --rm -v $(shell pwd)/api/openapi:/spec	redocly/cli join \
	/spec/demo.yaml \
	/spec/user.yaml \
	-o ./swagger.yaml

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi/generate_server.sh demo internal/demo/ports ports
	@./scripts/openapi/generate_server.sh user internal/user/ports ports

.PHONY: openapi_js
openapi_js:
	@./scripts/openapi/generate_client.sh demo
	@./scripts/openapi/generate_client.sh user
