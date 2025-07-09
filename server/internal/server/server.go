package server

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"

	pokemonpb "github.com/gimmy/awesome-api/gen/pokemon"
)

type pokemonServer struct {
	pokemonpb.UnimplementedPokemonServiceServer
}

func (s *pokemonServer) ListPokemons(ctx context.Context, req *pokemonpb.ListPokemonsRequest) (*pokemonpb.ListPokemonsResponse, error) {
	return &pokemonpb.ListPokemonsResponse{
		Pokemons: []*pokemonpb.Pokemon{
			{Name: "Pikachu", Id: 25},
			{Name: "Bulbasaur", Id: 1},
			{Name: "Charmander", Id: 4},
			{Name: "Squirtle", Id: 7},
			{Name: "Jigglypuff", Id: 39},
			{Name: "Meowth", Id: 52},
			{Name: "Psyduck", Id: 54},
			{Name: "Gengar", Id: 94},
			{Name: "Eevee", Id: 133},
			{Name: "Snorlax", Id: 143},
			{Name: "Mewtwo", Id: 150},
			{Name: "Mew", Id: 151},
			{Name: "Lucario", Id: 448},
			{Name: "Greninja", Id: 658},
			{Name: "Garchomp", Id: 445},
		},
	}, nil
}

func Run() error {
	grpcServer := grpc.NewServer()
	pokemonpb.RegisterPokemonServiceServer(grpcServer, &pokemonServer{})

	mux := runtime.NewServeMux()
	ctx := context.Background()
	if err := pokemonpb.RegisterPokemonServiceHandlerServer(ctx, mux, &pokemonServer{}); err != nil {
		return err
	}

	// Set up CORS middleware (allow all origins for demo)
	handler := cors.AllowAll().Handler(mux)

	go func() {
		log.Println("Serving gRPC on :9090")
		if err := http.ListenAndServe(":9090", grpcServer); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	log.Println("Serving HTTP on :8080 with CORS")
	return http.ListenAndServe(":8080", handler)
}
