package docx

type ReplaceDocx struct {
	zipReader ZipData
	content   string
	links     string
	headers   map[string]string
	footers   map[string]string
}

func (r *ReplaceDocx) Editable() *Docx {
	return &Docx{
		files:   r.zipReader.files(),
		content: r.content,
		links:   r.links,
		headers: r.headers,
		footers: r.footers,
	}
}

func (r *ReplaceDocx) Close() error {
	return r.zipReader.close()
}
