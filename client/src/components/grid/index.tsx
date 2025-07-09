import type { Pokemon } from "../../types/pokemons";

export default function Grid({ pokemon }: { pokemon: Pokemon }) {
  return (
    <div
      key={pokemon.id ?? Math.random()}
      style={{
        background: "white",
        borderRadius: 16,
        boxShadow: "0 4px 24px rgba(60,72,100,0.12)",
        padding: "1.5rem 1rem",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        transition: "transform 0.2s, box-shadow 0.2s",
        cursor: "pointer",
        border: "2px solid #e0e0e0",
        position: "relative",
      }}
      onMouseOver={(e) => {
        (e.currentTarget as HTMLDivElement).style.transform =
          "translateY(-6px) scale(1.04)";
        (e.currentTarget as HTMLDivElement).style.boxShadow =
          "0 8px 32px rgba(60,72,100,0.18)";
        (e.currentTarget as HTMLDivElement).style.border = "2px solid #3b4cca";
      }}
      onMouseOut={(e) => {
        (e.currentTarget as HTMLDivElement).style.transform = "";
        (e.currentTarget as HTMLDivElement).style.boxShadow =
          "0 4px 24px rgba(60,72,100,0.12)";
        (e.currentTarget as HTMLDivElement).style.border = "2px solid #e0e0e0";
      }}
    >
      {pokemon.name && (
        <img
          src={`https://img.pokemondb.net/artwork/avif/${pokemon.name.toLowerCase()}.avif`}
          alt={pokemon.name}
          style={{
            width: 120,
            height: 120,
            objectFit: "contain",
            marginBottom: 16,
            background: "#f6f8fc",
            borderRadius: 12,
            boxShadow: "0 2px 8px rgba(60,72,100,0.08)",
          }}
        />
      )}
      <div
        style={{
          fontWeight: 700,
          fontSize: "1.2rem",
          color: "#3b4cca",
          marginBottom: 4,
          letterSpacing: 1,
        }}
      >
        {pokemon.name}
      </div>
      <div style={{ color: "#888", fontSize: "1rem" }}>
        #{pokemon.id ? pokemon.id.toString().padStart(3, "0") : "???"}
      </div>
    </div>
  );
}
