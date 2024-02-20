package flag

import "flag"

type Flags struct {
	Tail  []string
	Bytes bool
	Lines bool
	Words bool
	Chars bool
}

func GetFlags() Flags {
	var byteFlag bool
	var lineFlag bool
	var wordFlag bool
	var charFlag bool

	flag.BoolVar(&byteFlag, "c", false, "Print byte counts")
	flag.BoolVar(&lineFlag, "l", false, "Print line counts")
	flag.BoolVar(&wordFlag, "w", false, "Print word counts")
	flag.BoolVar(&charFlag, "m", false, "Print character counts")

	flag.Parse()

	if !byteFlag && !lineFlag && !wordFlag && !charFlag {
		byteFlag = true
		lineFlag = true
		wordFlag = true
	}

	tail := flag.Args()

	return Flags{Bytes: byteFlag, Lines: lineFlag, Words: wordFlag, Chars: charFlag, Tail: tail}
}
