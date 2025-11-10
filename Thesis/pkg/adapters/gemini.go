package adapters
model string
embModel string
http *http.Client
}


func NewGemini(apiKey, model, embModel string) *Gemini {
if model == "" { model = "gemini-1.5-pro-latest" }
if embModel == "" { embModel = "text-embedding-004" }
return &Gemini{apiKey: apiKey, model: model, embModel: embModel, http: &http.Client{Timeout: 60 * time.Second}}
}


// LLM
var _ ports.LLM = (*Gemini)(nil)


// Embedder
var _ ports.Embedder = (*Gemini)(nil)


type part struct { Text string `json:"text"` }


type content struct { Role string `json:"role"`; Parts []part `json:"parts"` }


type genReq struct {
Contents []content `json:"contents"`
SystemInstruction *content `json:"systemInstruction,omitempty"`
}


type genResp struct {
Candidates []struct { Content struct { Parts []part `json:"parts"` } `json:"content"` } `json:"candidates"`
}


func (g *Gemini) Generate(system, user string) (string, error) {
url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", g.model, g.apiKey)
payload := genReq{Contents: []content{{Role: "user", Parts: []part{{Text: user}}}}, SystemInstruction: &content{Role:"system", Parts: []part{{Text: system}}}}
b, _ := json.Marshal(payload)
req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
req.Header.Set("Content-Type", "application/json")
res, err := g.http.Do(req)
if err != nil { return "", err }
defer res.Body.Close()
if res.StatusCode != 200 { x,_ := io.ReadAll(res.Body); return "", fmt.Errorf("gemini generate: %d %s", res.StatusCode, string(x)) }
var out genResp
if err := json.NewDecoder(res.Body).Decode(&out); err != nil { return "", err }
if len(out.Candidates)==0 || len(out.Candidates[0].Content.Parts)==0 { return "", fmt.Errorf("no candidates") }
return out.Candidates[0].Content.Parts[0].Text, nil
}


type embReq struct { Content struct { Parts []part `json:"parts"` } `json:"content"` }


type embResp struct { Embedding struct { Values []float64 `json:"values"` } `json:"embedding"` }


func (g *Gemini) Embed(text string) ([]float64, error) {
url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:embedContent?key=%s", g.embModel, g.apiKey)
var payload embReq
payload.Content.Parts = []part{{Text: text}}
b,_ := json.Marshal(payload)
req,_ := http.NewRequest("POST", url, bytes.NewReader(b))
req.Header.Set("Content-Type", "application/json")
res, err := g.http.Do(req)
if err != nil { return nil, err }
defer res.Body.Close()
if res.StatusCode != 200 { x,_ := io.ReadAll(res.Body); return nil, fmt.Errorf("gemini embed: %d %s", res.StatusCode, string(x)) }
var out embResp
if err := json.NewDecoder(res.Body).Decode(&out); err != nil { return nil, err }
return out.Embedding.Values, nil
}