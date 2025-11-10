package domain


import (
"encoding/json"
"fmt"
"strings"
"paperagent/pkg/ports"
)


type Summarizer struct {
LLM ports.LLM
Embedder ports.Embedder
Vector ports.VectorStore
}


type ctxItem struct { Text, Section, Heading string; Page int }


func (s *Summarizer) Summarize(collection, heading, systemPrompt string, topK int) (string, []map[string]any, error) {
if topK <= 0 { topK = 10 }
qvec, err := s.Embedder.Embed(heading + " の要点")
if err != nil { return "", nil, err }
ctxs, err := s.Vector.Search(collection, qvec, topK)
if err != nil { return "", nil, err }


// user プロンプト
items := make([]ctxItem, 0, len(ctxs))
for _, c := range ctxs {
it := ctxItem{}
if v,ok := c["text"].(string); ok { it.Text = v }
if v,ok := c["section"].(string); ok { it.Section = v }
if v,ok := c["heading"].(string); ok { it.Heading = v }
if v,ok := c["page"].(float64); ok { it.Page = int(v) }
items = append(items, it)
}
b, _ := json.MarshalIndent(items, "", " ")
user := fmt.Sprintf("# 要約対象セクション: %s\n\n# 原文抜粋(JSON)\n%s\n\n# 出力仕様\n- 箇条書き（最大5）\n- 各文末に (p.XX §YY)", strings.TrimSpace(heading), string(b))


out, err := s.LLM.Generate(systemPrompt, user)
return out, ctxs, err
}