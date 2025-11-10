package app


import (
_ "embed"
"os"
"path/filepath"
"paperagent/pkg/config"
)


type Prompts struct {
SectionSummary string
Plan string
QA string
}


//go:embed prompts/ja/summary_system.txt
var embedSummary string
//go:embed prompts/ja/plan_system.txt
var embedPlan string
//go:embed prompts/ja/qa_system.txt
var embedQA string


func LoadPrompts(cfg *config.Config) *Prompts {
p := &Prompts{SectionSummary: embedSummary, Plan: embedPlan, QA: embedQA}
if cfg.PromptsDir == "" { return p }
read := func(rel string) (string, bool) {
b, err := os.ReadFile(filepath.Join(cfg.PromptsDir, rel))
if err != nil || len(b) == 0 { return "", false }
return string(b), true
}
if s, ok := read(filepath.Join("ja","summary_system.txt")); ok { p.SectionSummary = s }
if s, ok := read(filepath.Join("ja","plan_system.txt")); ok { p.Plan = s }
if s, ok := read(filepath.Join("ja","qa_system.txt")); ok { p.QA = s }
return p
}