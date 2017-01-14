package main

type SmalltalkCounter struct {
	fileType string
}

func (u *SmalltalkCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{}, []MultiLineComment{MultiLineComment{begin: "\"", end: "\""}})
}
