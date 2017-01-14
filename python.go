package main

type PythonCounter struct {
	fileType string
}

func (u *PythonCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"#"}, []MultiLineComment{
		MultiLineComment{begin: "'''", end: "'''"},
		MultiLineComment{begin: "\"\"\"", end: "\"\"\""},
	})
}
