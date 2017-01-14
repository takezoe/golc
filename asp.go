package main

type ASPCounter struct {
	fileType string
}

func (u *ASPCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"'"}, []MultiLineComment{MultiLineComment{begin: "<!--", end: "-->"}})
}
