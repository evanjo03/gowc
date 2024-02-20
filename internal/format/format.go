package format

import (
	"fmt"

	wcfile "github.com/evanjo03/gowc/internal/file"
	wcflag "github.com/evanjo03/gowc/internal/flag"
)

func GetFileName(tail []string) string {
	if len(tail) == 0 {
		return ""
	}
	return tail[0]
}

func GetOutput(flags wcflag.Flags, fileCounts wcfile.File) string {
	var output string

	fileName := GetFileName(flags.Tail)

	if flags.Lines {
		output += fmt.Sprintf(" %d", fileCounts.Lines)
	}
	if flags.Words {
		output += fmt.Sprintf(" %d", fileCounts.Words)
	}
	if flags.Chars {
		output += fmt.Sprintf(" %d", fileCounts.Chars)
	}
	if flags.Bytes {
		output += fmt.Sprintf(" %d", fileCounts.Bytes)
	}
	if fileName != "" {
		output += fmt.Sprintf(" %s", fileName)
	}

	return output
}
