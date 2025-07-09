package server

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pokemonpb "github.com/gimmy/awesome-api/pokemon"
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

	go func() {
		log.Println("Serving gRPC on :9090")
		if err := http.ListenAndServe(":9090", grpcServer); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	log.Println("Serving HTTP on :8080")
	return http.ListenAndServe(":8080", mux)
}
