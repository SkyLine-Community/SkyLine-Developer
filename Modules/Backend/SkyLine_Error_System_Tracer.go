package SkyLine_Backend

import (
	"io/ioutil"
	"log"
	"strings"
)

func Output(char string, indentby int, increaseindent int) string {
	charliner := strings.Repeat(" ", indentby-1)
	charliner += strings.Repeat(" ", increaseindent) + char
	return charliner
}

// Scan last few lines within the code file given a certain line number
// For example if you get line 15 as an error the program will scan from the 10th line until the 15th line
// in other words report the last 5 lines before the error
func GrabLast5LinesBasedOnIntegerInputFromLineTracerErrorSystem(filename string, erroredline int) (Last5 []string, errorline string) {
	bytes, x := ioutil.ReadFile(filename)
	if x != nil {
		log.Fatal(x)
	}
	Content := string(bytes)
	lines := strings.Split(Content, "\n")
	if len(lines) <= 6 {
		return nil, ""
	}
	GetLastLineBeforeError := erroredline - 2
	for i := 0; i < 6; i++ {
		Last5 = append(Last5, string(lines[GetLastLineBeforeError-i]))
	}
	ReverseArrayForFileTraceback(Last5)
	return Last5, lines[erroredline-1]
}
