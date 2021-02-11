package docx

import "archive/zip"

//Contains functions to work with data from a zip file
type ZipData interface {
	files() []*zip.File
	close() error
}

//Type for in memory zip files
type ZipInMemory struct {
	data *zip.Reader
}

//Type for zip files read from disk
type ZipFile struct {
	data *zip.ReadCloser
}

func (d ZipInMemory) files() []*zip.File {
	return d.data.File
}

//Since there is nothing to close for in memory, just nil the data and return nil
func (d ZipInMemory) close() error {
	d.data = nil
	return nil
}

func (d ZipFile) files() []*zip.File {
	return d.data.File
}

func (d ZipFile) close() error {
	return d.data.Close()
}
