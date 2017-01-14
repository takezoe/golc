package main

type XMLCounter struct {
	fileType string
}

func (u *XMLCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{}, []MultiLineComment{MultiLineComment{begin: "<!--", end: "-->"}})
}
