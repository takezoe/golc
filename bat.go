package main

type BatCounter struct {
	fileType string
}

func (u *BatCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"REM"}, []MultiLineComment{})
}
