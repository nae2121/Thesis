package config


import (
"os"
"strconv"
)


type Config struct {
Port string
GoogleAPIKey string
GeminiModel string
EmbeddingModel string
GrobidURL string
QdrantURL string
EmbeddingDim int
DataDir string
PromptsDir string
}


func Load() *Config {
return &Config{
Port: getenv("PORT", "8080"),
GoogleAPIKey: mustenv("GOOGLE_API_KEY"),
GeminiModel: getenv("GEMINI_MODEL", "gemini-1.5-pro-latest"),
EmbeddingModel: getenv("GEMINI_EMBEDDING_MODEL", "text-embedding-004"),
GrobidURL: getenv("GROBID_URL", "http://grobid:8070"),
QdrantURL: getenv("QDRANT_URL", "http://qdrant:6333"),
EmbeddingDim: getint("EMBEDDING_DIM", 768),
DataDir: getenv("DATA_DIR", "/app/data"),
PromptsDir: getenv("PROMPTS_DIR", ""),
}
}


func getenv(k, def string) string { if v := os.Getenv(k); v != "" { return v }; return def }
func mustenv(k string) string { v := os.Getenv(k); if v == "" { panic("missing env: "+k) }; return v }
func getint(k string, def int) int { if v := os.Getenv(k); v != "" { if i,err:=strconv.Atoi(v); err==nil { return i } }; return def }