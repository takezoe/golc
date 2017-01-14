package main

type PHPCounter struct {
	fileType string
}

func (u *PHPCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"//"}, []MultiLineComment{
		MultiLineComment{begin: "/*", end: "*/"},
		MultiLineComment{begin: "<!--", end: "-->"},
	})
}
