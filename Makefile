include .env

.PHONY: bootstrap
bootstrap:  openapi_http

.PHONY: openapi_http
openapi_http:
	@tput setaf 2;echo "generate http server"
	@./scripts/openapi/generate_server.sh demo internal/demo/ports ports
	@./scripts/openapi/generate_server.sh user internal/user/ports ports
