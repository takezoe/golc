package main

type LuaCounter struct {
	fileType string
}

func (u *LuaCounter) Count(filePath string) CountResult {
	return countFile(filePath, u.fileType, []string{"--"}, []MultiLineComment{
		MultiLineComment{begin: "--[[", end: "]]"},
		MultiLineComment{begin: "--[===[", end: "]===]"},
	})
}
