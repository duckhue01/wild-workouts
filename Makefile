include .env

.PHONY: bootstrap
bootstrap: build_oapi_codegen openapi

.PHONY: openapi
openapi: openapi_http openapi_js

.PHONY: build_oapi_codegen
build_oapi_codegen:
	@tput setaf 2;echo "build oapi-codegen image"
	@docker build -f docker/oapi-codegen/Dockerfile -t oapi-codegen .

.PHONY: openapi_http
openapi_http:
	@tput setaf 2;echo "generate http server"
	@./scripts/openapi/generate_server.sh demo internal/demo/ports ports
	@./scripts/openapi/generate_server.sh user internal/user/ports ports
