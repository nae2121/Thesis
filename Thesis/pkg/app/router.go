package app


// embed & upsert
vecs := make([][]float64, 0, len(chunks))
payloads := make([]map[string]any, 0, len(chunks))
for _, ch := range chunks {
v, err := s.Embedder.Embed(ch.Text)
if err != nil { c.JSON(502, gin.H{"error": err.Error()}); return }
vecs = append(vecs, v)
payloads = append(payloads, map[string]any{"text": ch.Text, "page": ch.Page, "section": ch.Section, "heading": ch.Heading})
}
if err := s.Vector.Upsert(coll, vecs, payloads); err != nil { c.JSON(502, gin.H{"error": err.Error()}); return }


c.JSON(200, gin.H{"collection": coll, "chunks": len(chunks), "title": title})
})


type summarizeReq struct { Collection string `json:"collection"`; Heading string `json:"heading"`; TopK int `json:"top_k"` }
r.POST("/summarize", func(c *gin.Context) {
var req summarizeReq
if err := c.ShouldBindJSON(&req); err != nil { c.JSON(400, gin.H{"error": err.Error()}); return }
if req.Collection=="" { c.JSON(400, gin.H{"error":"collection is required"}); return }
if req.TopK<=0 { req.TopK=10 }


qvec, err := s.Embedder.Embed(req.Heading+" の要点")
if err != nil { c.JSON(502, gin.H{"error": err.Error()}); return }
ctxs, err := s.Vector.Search(req.Collection, qvec, req.TopK)
if err != nil { c.JSON(502, gin.H{"error": err.Error()}); return }


type ctxItem struct { Text, Section, Heading string; Page int }
items := make([]ctxItem, 0, len(ctxs))
for _, v := range ctxs {
it := ctxItem{}
if t,ok:=v["text"].(string); ok { it.Text=t }
if sct,ok:=v["section"].(string); ok { it.Section=sct }
if hd,ok:=v["heading"].(string); ok { it.Heading=hd }
if pg,ok:=v["page"].(float64); ok { it.Page=int(pg) }
items = append(items, it)
}
b, _ := json.MarshalIndent(items, "", " ")
user := fmt.Sprintf("# 要約対象セクション: %s\n\n# 原文抜粋(JSON)\n%s\n\n# 出力仕様\n- 箇条書き（最大5）\n- 各文末に (p.XX §YY)", req.Heading, string(b))


out, err := s.LLM.Generate(s.Prompts.SectionSummary, user)
if err != nil { c.JSON(502, gin.H{"error": err.Error()}); return }
c.JSON(200, gin.H{"heading": req.Heading, "summary": out, "evidence": ctxs})
})
}