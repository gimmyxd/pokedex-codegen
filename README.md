# Pokédex

- **Code Generation:** Automated code generation.

## Project Structure

```
pokedex/                    # Project root
├── proto/                  # Protobuf definitions (PokemonService, etc.)
├── server/                 # Go service implementation
├── main.go                 # Backend entrypoint
├── Makefile                # Automation (generate, cleanup)
├── buf.gen.yaml            # Buf codegen config
├── buf.work.yaml           # Buf workspace config
├── client/                 # React frontend
│   ├── src/
│   │   ├── api/            # Orval-generated API hooks and types
│   │   ├── components/     # UI components (e.g., Grid)
│   │   ├── App.tsx         # Main app UI
│   │   └── App.css         # Global styles
│   ├── orval.config.ts     # Orval config for codegen
│   ├── package.json        # Frontend dependencies/scripts
│   └── README.md           # Frontend usage and codegen docs
└── README.md               # (this file)
```

## How it Works
- **Schema first:** Define your API in `proto/`.
- **Buf:** Generates Go code, gRPC server, grpc-gateway, and OpenAPI schema.
- **Go Backend:** Implements the PokemonService, serves both gRPC and REST endpoints
- **React Frontend:** Uses Orval to generate type-safe hooks and models from the OpenAPI schema, providing a beautiful, interactive Pokédex UI.

See each subdirectory's README for more details on setup, code generation, and running the app.
