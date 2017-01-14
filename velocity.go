package main

type VelocityCounter struct {
	fileType string
}

func (u *VelocityCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"##"}, []MultiLineComment{})
}
