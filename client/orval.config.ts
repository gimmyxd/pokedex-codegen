import { defineConfig } from 'orval';

export default defineConfig({
  pokemons: {
    input: {
      target: "../server/gen/openapi/openapi.json"
    },
    output: {
      clean: false,
      client: 'react-query',
      mock: true,
      mode: 'split',
      schemas: 'src/types/pokemons',
      target: 'src/api/v1/pokemons.ts',
      baseUrl: 'http://localhost:8080',
    },
  },
})
