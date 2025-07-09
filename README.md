# Awesome API: Pokemon Service

This project demonstrates a Go gRPC service for listing Pokemons, with a RESTful HTTP gateway and OpenAPI (Swagger) schema generation using [Buf](https://buf.build/) and remote plugins.

## Features
- gRPC API for listing Pokemons
- RESTful HTTP/JSON API via grpc-gateway
- OpenAPI (Swagger) schema generation (YAML & JSON)
- Modern code generation with Buf and remote plugins

## Project Structure
```
proto/           # Protobuf definitions
pokemon/         # Generated Go code (protos, gRPC, gateway)
openapi/       # Generated OpenAPI schema (YAML & JSON)
internal/server/ # Service implementation
main.go          # Entrypoint
Makefile         # Automation scripts
buf.gen.yaml     # Buf codegen config
buf.work.yaml    # Buf workspace config
```

## Prerequisites
- Go 1.20+
- [Buf CLI](https://docs.buf.build/installation)
- [yq](https://github.com/mikefarah/yq) (for YAML→JSON conversion)

## Setup & Code Generation
```sh
make generate
```
This will:
- Generate Go code for protos, gRPC, and grpc-gateway
- Generate OpenAPI schema in both YAML and JSON formats

## Cleanup Generated Files
```sh
make cleanup
```

## Running the Service
```sh
go run main.go
```
- gRPC server: `localhost:9090`
- HTTP REST API: `localhost:8080`

## Example REST API Usage
List Pokemons:
```sh
curl http://localhost:8080/v1/pokemons
```
Response:
```json
{
  "pokemons": [
    {"name": "Pikachu", "id": 25},
    {"name": "Bulbasaur", "id": 1},
    {"name": "Charmander", "id": 4}
  ]
}
```

## OpenAPI Schema
- YAML: `openapi/openapi.yaml`
- JSON: `openapi/openapi.json`

You can use these with Swagger UI, Postman, or for client code generation.

## References
- [Buf CLI Quickstart](https://buf.build/docs/cli/quickstart/)
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [google-gnostic-openapi plugin](https://buf.build/community/google-gnostic-openapi)
