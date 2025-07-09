BUF_GEN_DIRS=gen

.PHONY: generate cleanup

generate:
	buf generate && yq -o=json gen/openapi/openapi.yaml > gen/openapi/openapi.json

cleanup:
	rm -rf $(BUF_GEN_DIRS)
