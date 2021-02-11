package docx

import "strings"

type Table struct {
	content     string
	templateRow string
	newRows     []string
}

func (t *Table) AddRow(row string) {
	t.newRows = append(t.newRows, row)
}

func (t *Table) GetTemplate() string {
	return t.templateRow
}

func (t *Table) generateNewTable() string {
	return strings.Replace(t.content, t.templateRow, strings.Join(t.newRows, ""), -1)
}

func getTemplateRow(content string) string {
	start := strings.LastIndex(content, "<w:tr>")
	end := strings.LastIndex(content, "</w:tr>")
	if start == -1 || end == -1 {
		return ""
	}
	end += len("</w:tr>")
	return content[start:end]
}

func getTableContent(content string) (res []string) {
	start := strings.Index(content, "<w:tbl>")
	end := strings.Index(content, "</w:tbl>")
	if start == -1 || end == -1 {
		return
	}
	end += len("</w:tbl>")
	res = []string{content[start:end]}
	tables := getTableContent(content[end:])
	for _, table := range tables {
		res = append(res, table)
	}
	return
}
