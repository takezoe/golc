package main

type VBCounter struct {
	fileType string
}

func (u *VBCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"'", "REM"}, []MultiLineComment{})
}
