package main

type JSPCounter struct {
	fileType string
}

func (u *JSPCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"//"}, []MultiLineComment{
		MultiLineComment{begin: "/*", end: "*/"},
		MultiLineComment{begin: "<%--", end: "--%>"},
		MultiLineComment{begin: "<!--", end: "-->"},
	})
}
