package app


import (
"fmt"
"paperagent/pkg/adapters"
"paperagent/pkg/config"
"paperagent/pkg/httpserver"
)


type Server struct {
Cfg *config.Config
Router interface{ Run(addr ...string) error }
// 依存（portsで抽象化）
LLM interface{ Generate(system, user string) (string, error) }
Embedder interface{ Embed(text string) ([]float64, error) }
Vector interface{ EnsureCollection(string) error; Upsert(string, [][]float64, []map[string]any) error; Search(string, []float64, int) ([]map[string]any, error) }
Parser interface{ ProcessFulltext(string) (string, error) }
Prompts *Prompts
}


func New() (*Server, error) {
cfg := config.Load()
r := httpserver.New()


g := adapters.NewGemini(cfg.GoogleAPIKey, cfg.GeminiModel, cfg.EmbeddingModel)
q := adapters.NewQdrant(cfg.QdrantURL, cfg.EmbeddingDim)
gr := adapters.NewGrobid(cfg.GrobidURL)


s := &Server{Cfg: cfg, Router: r, LLM: g, Embedder: g, Vector: q, Parser: gr, Prompts: LoadPrompts(cfg)}
registerRoutes(s)
return s, nil
}


func (s *Server) Run() error { return s.Router.Run(":" + s.Cfg.Port) }


// 依存関係が足りない場合の compile error を避けるダミー参照
var _ = fmt.Sprintf