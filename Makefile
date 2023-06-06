include .env

.PHONY: bootstrap
bootstrap: openapi

.PHONY: openapi
openapi:
	@tput setaf 2;echo "generate http server & http types"
	@./scripts/openapi/generate_server.sh demo internal/demo/ports ports
	@./scripts/openapi/generate_server.sh auth internal/auth main
	@./scripts/openapi/generate_server.sh notif internal/notif/ports ports


.PHONY: previewdoc
previewdoc:
	@tput setaf 2;echo "preview api"
	redocly join ./api/openapi/*.yaml -o ./main.yaml
	redocly preview-docs ./main.yaml

.PHONY: eventproto
eventproto:
	@tput setaf 2;echo "generating event proto"
	buf generate --template api/proto/events/buf.gen.yaml

.PHONY: mergedoc
mergedoc:
	@tput setaf 2;echo "merge openapi docs"
	redocly join ./api/openapi/*.yaml -o ./main.yaml