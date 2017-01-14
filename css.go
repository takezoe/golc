package main

type CSSCounter struct {
	fileType string
}

func (u *CSSCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{}, []MultiLineComment{MultiLineComment{begin: "/*", end: "*/"}})
}
