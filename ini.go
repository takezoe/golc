package main

type IniCounter struct {
	fileType string
}

func (u *IniCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{";"}, []MultiLineComment{})
}
