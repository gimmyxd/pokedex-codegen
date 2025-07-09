import "./App.css";
import { usePokemonServiceListPokemons } from "./api/v1/pokemons";
import Grid from "./components/grid";

function App() {
  const { data } = usePokemonServiceListPokemons();
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        minHeight: "100vh",
        background: "linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%)",
        padding: "2rem",
      }}
    >
      <h1
        style={{
          fontFamily: "Pokemon, sans-serif",
          color: "#ffcb05",
          textShadow: "2px 2px #3b4cca",
          marginBottom: "2rem",
          fontSize: "2.5rem",
          letterSpacing: 2,
        }}
      >
        Pokédex
      </h1>
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "repeat(auto-fit, minmax(220px, 1fr))",
          gap: "2rem",
          width: "100%",
          maxWidth: 900,
        }}
      >
        {data?.data.pokemons?.map((pokemon) => (
          <Grid pokemon={pokemon} />
        ))}
      </div>
    </div>
  );
}

export default App;
