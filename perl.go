package main

type PerlCounter struct {
	fileType string
}

func (u *PerlCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"#"}, []MultiLineComment{MultiLineComment{begin: "=pod", end: "=cut"}})
}
