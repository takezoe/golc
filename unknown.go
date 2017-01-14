package main

type UnknownCounter struct{}

func (u *UnknownCounter) Count(filePath string) CountResult {
	return CountResult{FilePath: filePath, FileType: "Unknown", Code: 0, Empty: 0, Comment: 0}
}
