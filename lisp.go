package main

type LispCounter struct {
	fileType string
}

func (u *LispCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{";"}, []MultiLineComment{})
}
