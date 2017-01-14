package main

type ShellCounter struct {
	fileType string
}

func (u *ShellCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"#"}, []MultiLineComment{})
}
