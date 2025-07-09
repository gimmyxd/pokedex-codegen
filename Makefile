BUF_GEN_DIRS=server/gen

.PHONY: generate cleanup

generate:
	buf generate && yq -o=json server/gen/openapi/openapi.yaml > server/gen/openapi/openapi.json

cleanup:
	rm -rf $(BUF_GEN_DIRS)
