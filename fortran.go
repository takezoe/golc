package main

type FortranCounter struct {
	fileType string
}

func (u *FortranCounter) Count(filePath string) CountResult {
	// TODO *, !, c at the head of line
	return countFile(filePath, u.fileType, []string{"*", "!"}, []MultiLineComment{})
}
