package domain


import (
"encoding/xml"
"strings"
)


type teiTEI struct {
XMLName xml.Name `xml:"TEI"`
Title string `xml:"teiHeader>fileDesc>titleStmt>title"`
Text struct{ Body teiBody `xml:"body"` } `xml:"text"`
}


type teiBody struct { Divs []teiDiv `xml:"div"` }


type teiDiv struct {
Head string `xml:"head"`
Ps []string `xml:"p"`
Divs []teiDiv `xml:"div"`
}


func ChunkFromTEI(teiXML string, maxLen int) ([]Chunk, string, error) {
if maxLen <= 0 { maxLen = 1200 }
var t teiTEI
if err := xml.Unmarshal([]byte(teiXML), &t); err != nil { return nil, "", err }
var out []Chunk
var walk func(d teiDiv, prefix string)
walk = func(d teiDiv, prefix string) {
sec := prefix
if strings.TrimSpace(d.Head) != "" {
if sec != "" { sec += "." }
sec += oneLine(d.Head)
}
text := strings.Join(d.Ps, "\n")
for _, seg := range splitRunes(oneLine(text), maxLen) {
if strings.TrimSpace(seg) == "" { continue }
out = append(out, Chunk{Text: seg, Page: -1, Section: sec, Heading: oneLine(d.Head)})
}
for _, ch := range d.Divs { walk(ch, sec) }
}
for _, d := range t.Text.Body.Divs { walk(d, "") }
return out, oneLine(t.Title), nil
}


func oneLine(s string) string { s=strings.ReplaceAll(s,"\n"," "); return strings.Join(strings.Fields(s)," ") }


func splitRunes(s string, max int) []string {
r := []rune(strings.TrimSpace(s))
if len(r) <= max { return []string{string(r)} }
var parts []string
for i:=0; i<len(r); i+=max { j:=i+max; if j>len(r){j=len(r)}; parts=append(parts, string(r[i:j])) }
return parts
}