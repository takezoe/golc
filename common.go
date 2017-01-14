package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type CountResult struct {
	FilePath string
	FileType string
	Code     int
	Empty    int
	Comment  int
}

type MultiLineComment struct {
	begin string
	end   string
}

type Counter interface {
	Count(filePath string) CountResult
}

func countFile(filePath string, fileType string, singleComments []string, multiLineComments []MultiLineComment) CountResult {
	fp, _ := os.Open(filePath)
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 1024*8)

	result := CountResult{FilePath: filePath, FileType: fileType, Code: 0, Empty: 0, Comment: 0}
	inMultiLineComment := false

	for {
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		line := strings.TrimSpace(string(bytes))
		var endOfMultiComment string

		if inMultiLineComment == false {
			if len(line) == 0 {
				// empty line
				result.Empty = result.Empty + 1

			} else if checkSingleLineComment(line, singleComments) {
				// single line comment
				result.Comment = result.Comment + 1

			} else if endOfMultiComment = checkMultiLineComment(line, multiLineComments); endOfMultiComment != "" {
				// begin multi line comment
				inMultiLineComment = true
				result.Comment = result.Comment + 1

			} else {
				// effective line
				result.Code = result.Code + 1
			}
		} else {
			result.Comment = result.Comment + 1
			if strings.Contains(line, endOfMultiComment) {
				// end of multi line comment
				inMultiLineComment = false
			}
		}
	}

	return result
}

func checkSingleLineComment(line string, singleLineComments []string) bool {
	for _, prefix := range singleLineComments {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

func checkMultiLineComment(line string, multiLineComments []MultiLineComment) string {
	for _, multiLineComment := range multiLineComments {
		index := strings.Index(line, multiLineComment.begin)
		if index >= 0 && strings.Index(line[index:len(line)], multiLineComment.end) < 0 {
			return multiLineComment.end
		}
	}
	return ""
}

func getCounter(filePath string) Counter {
	if strings.HasSuffix(filePath, ".java") {
		return &JavaCounter{fileType: "Java"}

	} else if strings.HasSuffix(filePath, ".scala") {
		return &JavaCounter{fileType: "Scala"}

	} else if strings.HasSuffix(filePath, ".cpp") || strings.HasSuffix(filePath, ".cxx") || strings.HasSuffix(filePath, ".cc") || strings.HasSuffix(filePath, ".c") {
		return &JavaCounter{fileType: "C/C++/ObjC"}

	} else if strings.HasSuffix(filePath, ".h") || strings.HasSuffix(filePath, ".hpp") || strings.HasSuffix(filePath, ".hxx") {
		return &JavaCounter{fileType: "h"}

	} else if strings.HasSuffix(filePath, ".cs") {
		return &JavaCounter{fileType: "C#"}

	} else if strings.HasSuffix(filePath, ".js") {
		return &JavaCounter{fileType: "JavaScript"}

	} else if strings.HasSuffix(filePath, ".json") {
		return &JavaCounter{fileType: "JSON"} // TODO JSON doesn't contain comment

	} else if strings.HasSuffix(filePath, ".vbs") {
		return &VBCounter{fileType: "VBScript"}

	} else if strings.HasSuffix(filePath, ".bas") || strings.HasSuffix(filePath, ".frm") || strings.HasSuffix(filePath, ".cls") {
		return &VBCounter{fileType: "VB"}

	} else if strings.HasSuffix(filePath, ".vb") {
		return &VBCounter{fileType: "VB.NET"}

	} else if strings.HasSuffix(filePath, ".go") {
		return &JavaCounter{fileType: "Go"}

	} else if strings.HasSuffix(filePath, ".as") {
		return &JavaCounter{fileType: "ActionScript"}

	} else if strings.HasSuffix(filePath, ".grooby") {
		return &JavaCounter{fileType: "Groovy"}

	} else if strings.HasSuffix(filePath, ".html") {
		return &XMLCounter{fileType: "HTML"}

	} else if strings.HasSuffix(filePath, ".xhtml") {
		return &XMLCounter{fileType: "XHTML"}

	} else if strings.HasSuffix(filePath, ".xml") || strings.HasSuffix(filePath, ".dicon") {
		return &XMLCounter{fileType: "XML"}

	} else if strings.HasSuffix(filePath, ".dtd") {
		return &XMLCounter{fileType: "DTD"}

	} else if strings.HasSuffix(filePath, ".xsd") {
		return &XMLCounter{fileType: "XMLSchema"}

	} else if strings.HasSuffix(filePath, ".tld") {
		return &XMLCounter{fileType: "TLD"}

	} else if strings.HasSuffix(filePath, ".xsl") {
		return &XMLCounter{fileType: "XSLT"}

	} else if strings.HasSuffix(filePath, ".mxml") {
		return &XMLCounter{fileType: "MXML"}

	} else if strings.HasSuffix(filePath, ".xi") {
		return &XMLCounter{fileType: "Xi"}

	} else if strings.HasSuffix(filePath, ".tcl") {
		return &ShellCounter{fileType: "Tcl"}

	} else if strings.HasSuffix(filePath, ".sh") {
		return &ShellCounter{fileType: "Shell"}

	} else if strings.HasSuffix(filePath, ".properties") {
		return &ShellCounter{fileType: "Properties"}

		// } else if strings.HasSuffix(filePath, "Makefile") { // TODO get filename
		// 	return &ShellCounter{fileType: "Makefile"}

	} else if strings.HasSuffix(filePath, ".l") || strings.HasSuffix(filePath, ".el") || strings.HasSuffix(filePath, ".cl") {
		return &LispCounter{fileType: "Lisp"}

	} else if strings.HasSuffix(filePath, ".clj") {
		return &LispCounter{fileType: "Clojure"}

	} else if strings.HasSuffix(filePath, ".scm") {
		return &LispCounter{fileType: "Scheme"}

	} else if strings.HasSuffix(filePath, ".st") {
		return &SmalltalkCounter{fileType: "Smalltalk"}

	} else if strings.HasSuffix(filePath, ".jsp") {
		return &JSPCounter{fileType: "JSP"}

	} else if strings.HasSuffix(filePath, ".php") || strings.HasSuffix(filePath, ".php3") {
		return &PHPCounter{fileType: "PHP"}

	} else if strings.HasSuffix(filePath, ".asp") || strings.HasSuffix(filePath, ".asa") {
		return &ASPCounter{fileType: "ASP"}

	} else if strings.HasSuffix(filePath, ".cfm") {
		return &ColdFusionCounter{fileType: "ColdFusion"}

	} else if strings.HasSuffix(filePath, ".pl") || strings.HasSuffix(filePath, ".pm") {
		return &PerlCounter{fileType: "Perl"}

	} else if strings.HasSuffix(filePath, ".rb") {
		return &RubyCounter{fileType: "Ruby"}

	} else if strings.HasSuffix(filePath, ".sql") {
		return &SQLCounter{fileType: "SQL"}

	} else if strings.HasSuffix(filePath, ".bat") {
		return &BatCounter{fileType: "BAT"}

	} else if strings.HasSuffix(filePath, ".lua") {
		return &LuaCounter{fileType: "Lua"}

	} else if strings.HasSuffix(filePath, ".hs") {
		return &HaskellCounter{fileType: "Haskell"}

	} else if strings.HasSuffix(filePath, ".f") || strings.HasSuffix(filePath, ".for") || strings.HasSuffix(filePath, ".ftn") ||
		strings.HasSuffix(filePath, ".f90") || strings.HasSuffix(filePath, ".f95") {
		return &FortranCounter{fileType: "Fortran"}

	} else if strings.HasSuffix(filePath, ".ini") {
		return &IniCounter{fileType: "INI"}

	} else if strings.HasSuffix(filePath, ".vm") {
		return &VelocityCounter{fileType: "Velocity"}

	} else if strings.HasSuffix(filePath, ".py") {
		return &PythonCounter{fileType: "Python"}

	} else {
		return &UnknownCounter{}
	}
}
