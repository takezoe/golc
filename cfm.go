package main

type ColdFusionCounter struct {
	fileType string
}

func (u *ColdFusionCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{}, []MultiLineComment{
		MultiLineComment{begin: "<!--", end: "-->"},
		MultiLineComment{begin: "<!---", end: "--->"},
	})
}
