# Pokédex

Type-safe API hooks and types generated with [Orval](https://orval.dev/)

## Project Structure
```
src/
  api/           # Orval-generated API hooks and types
  components/    # UI components (e.g., Grid)
  App.tsx        # Main app UI
  App.css        # Global styles
```

## Setup
1. Install dependencies:
   ```sh
   yarn install
   ```

2. Generate API hooks/types from OpenAPI schema using Orval:
   ```sh
   yarn generate
   ```
   - This uses the OpenAPI schema (e.g., `../openapi/openapi.json`) to generate React Query hooks and TypeScript types in `src/api/`.
   - See [Orval documentation](https://orval.dev/docs/cli/overview) for more options.

3. Start the development server:
   ```sh
   yarn dev
   ```

## Code Generation with Orval
- Orval reads your OpenAPI schema and generates:
  - TypeScript types for all API models
  - React Query hooks for all endpoints
- The config file (`orval.config.ts`) specifies the OpenAPI input and output locations.
- Run `yarn generate` whenever the OpenAPI schema changes.

## Example Usage
```tsx
import { usePokemonServiceListPokemons } from './api/v1/pokemons';

const { data } = usePokemonServiceListPokemons();
```

## References
- [Orval Documentation](https://orval.dev/)
- [Pokédex API Backend](../README.md)
