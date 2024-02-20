package file

import (
	"bufio"
	"flag"
	"io"
	"os"
	"unicode"
)

type File struct {
	Bytes int64
	Lines int64
	Words int64
	Chars int64
}

func GetReader() (*bufio.Reader, error) {
	fileNames := flag.Args()

	if len(fileNames) == 0 {
		return bufio.NewReader(os.Stdin), nil
	}

	f, err := os.Open(fileNames[0])
	if err != nil {
		return nil, err
	}
	return bufio.NewReader(f), nil
}

func GetCounts(reader *bufio.Reader) File {
	var prevRune rune
	var bytesCount int64 = 0
	var linesCount int64 = 0
	var wordsCount int64 = 0
	var charsCount int64 = 0

	for {
		rune, size, err := reader.ReadRune()
		bytesCount += int64(size)

		if err == io.EOF {
			break
		}

		charsCount++

		if rune == '\n' {
			linesCount += 1
		}

		if !unicode.IsSpace(prevRune) && unicode.IsSpace(rune) {
			wordsCount += 1
		}

		prevRune = rune

	}

	return File{
		Bytes: bytesCount,
		Lines: linesCount,
		Words: wordsCount,
		Chars: charsCount,
	}
}
