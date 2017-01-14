package main

type HaskellCounter struct {
	fileType string
}

func (u *HaskellCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"--"}, []MultiLineComment{MultiLineComment{begin: "{-", end: "-}"}})
}
