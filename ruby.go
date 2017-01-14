package main

type RubyCounter struct {
	fileType string
}

func (u *RubyCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"#"}, []MultiLineComment{MultiLineComment{begin: "=begin", end: "=end"}})
}
