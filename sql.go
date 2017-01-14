package main

type SQLCounter struct {
	fileType string
}

func (u *SQLCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"#", "--", "REM"}, []MultiLineComment{MultiLineComment{begin: "/*", end: "*/"}})
}
