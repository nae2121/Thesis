package ports


type PaperParser interface {
	ProcessFulltext(pdfPath string) (string, error) // returns TEI XML
}