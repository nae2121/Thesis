package adapters


import (
"bytes"
"fmt"
"io"
"mime/multipart"
"net/http"
"os"
"time"
"paperagent/pkg/ports"
)


type Grobid struct { base string; http *http.Client }


func NewGrobid(base string) *Grobid {
if base == "" { base = "http://grobid:8070" }
return &Grobid{base: base, http: &http.Client{Timeout: 120 * time.Second}}
}


var _ ports.PaperParser = (*Grobid)(nil)


func (g *Grobid) ProcessFulltext(pdfPath string) (string, error) {
f, err := os.Open(pdfPath)
if err != nil { return "", err }
defer f.Close()
var buf bytes.Buffer
mw := multipart.NewWriter(&buf)
fw, _ := mw.CreateFormFile("input", pdfPath)
if _, err := io.Copy(fw, f); err != nil { return "", err }
_ = mw.WriteField("consolidateHeader", "1")
_ = mw.WriteField("consolidateCitations", "0")
_ = mw.WriteField("teiCoordinates", "false")
mw.Close()
req,_ := http.NewRequest("POST", g.base+"/api/processFulltextDocument", &buf)
req.Header.Set("Content-Type", mw.FormDataContentType())
res, err := g.http.Do(req)
if err != nil { return "", err }
defer res.Body.Close()
if res.StatusCode != 200 { x,_ := io.ReadAll(res.Body); return "", fmt.Errorf("grobid: %d %s", res.StatusCode, string(x)) }
b, _ := io.ReadAll(res.Body)
return string(b), nil
}