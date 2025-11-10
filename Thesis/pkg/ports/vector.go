package ports


type VectorStore interface {
	EnsureCollection(name string) error
	Upsert(collection string, vectors [][]float64, payloads []map[string]any) error
	Search(collection string, queryVec []float64, topK int) ([]map[string]any, error)
}