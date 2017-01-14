package main

type JavaCounter struct {
	fileType string
}

func (u *JavaCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"//"}, []MultiLineComment{MultiLineComment{begin: "/*", end: "*/"}})
}
