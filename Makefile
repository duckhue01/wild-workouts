include .env

.PHONY: openapi
openapi: openapi_http openapi_js merge_openapi_docs

.PHONY: merge_openapi_docs
merge_openapi_docs:
	docker run --rm -v /Users/duckhue01/code/side/wild-workouts/api/openapi:/spec	redocly/cli join \
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


.PHONY: sqlc
sqlc:
	sqlc generate