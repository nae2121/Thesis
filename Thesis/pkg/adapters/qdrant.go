package adapters
http *http.Client
}


func NewQdrant(base string, dim int) *Qdrant {
if base == "" { base = "http://qdrant:6333" }
if dim <= 0 { dim = 768 }
return &Qdrant{base: base, dim: dim, http: &http.Client{Timeout: 30 * time.Second}}
}


var _ ports.VectorStore = (*Qdrant)(nil)


type ensureReq struct { Vectors map[string]any `json:"vectors"` }


func (q *Qdrant) EnsureCollection(name string) error {
url := fmt.Sprintf("%s/collections/%s", q.base, name)
body := ensureReq{Vectors: map[string]any{"size": q.dim, "distance": "Cosine"}}
b,_ := json.Marshal(body)
req,_ := http.NewRequest("PUT", url, bytes.NewReader(b))
req.Header.Set("Content-Type","application/json")
res, err := q.http.Do(req)
if err != nil { return err }
defer res.Body.Close()
if res.StatusCode >= 400 { x,_ := io.ReadAll(res.Body); return fmt.Errorf("qdrant ensure: %s", string(x)) }
return nil
}


type upsertReq struct { Points []struct{ ID string `json:"id"`; Vector []float64 `json:"vector"`; Payload map[string]any `json:"payload"` } `json:"points"` }


func (q *Qdrant) Upsert(collection string, vectors [][]float64, payloads []map[string]any) error {
if len(vectors) != len(payloads) { return fmt.Errorf("vectors/payloads size mismatch") }
reqBody := upsertReq{Points: make([]struct{ID string; Vector []float64; Payload map[string]any}, len(vectors))}
for i := range vectors {
reqBody.Points[i].ID = fmt.Sprintf("%d", i)
reqBody.Points[i].Vector = vectors[i]
reqBody.Points[i].Payload = payloads[i]
}
b,_ := json.Marshal(reqBody)
url := fmt.Sprintf("%s/collections/%s/points?wait=true", q.base, collection)
req,_ := http.NewRequest("PUT", url, bytes.NewReader(b))
req.Header.Set("Content-Type","application/json")
res, err := q.http.Do(req)
if err != nil { return err }
defer res.Body.Close()
if res.StatusCode >= 300 { x,_ := io.ReadAll(res.Body); return fmt.Errorf("qdrant upsert: %s", string(x)) }
return nil
}


type searchResp struct { Result []struct{ Score float64 `json:"score"`; Payload map[string]any `json:"payload"` } `json:"result"` }


func (q *Qdrant) Search(collection string, queryVec []float64, topK int) ([]map[string]any, error) {
body := map[string]any{"vector": queryVec, "limit": topK, "with_payload": true}
b,_ := json.Marshal(body)
url := fmt.Sprintf("%s/collections/%s/points/search", q.base, collection)
req,_ := http.NewRequest("POST", url, bytes.NewReader(b))
req.Header.Set("Content-Type","application/json")
res, err := q.http.Do(req)
if err != nil { return nil, err }
defer res.Body.Close()
if res.StatusCode >= 300 { x,_ := io.ReadAll(res.Body); return nil, fmt.Errorf("qdrant search: %s", string(x)) }
var out searchResp
if err := json.NewDecoder(res.Body).Decode(&out); err != nil { return nil, err }
docs := make([]map[string]any, 0, len(out.Result))
for _, r := range out.Result { docs = append(docs, r.Payload) }
return docs, nil
}