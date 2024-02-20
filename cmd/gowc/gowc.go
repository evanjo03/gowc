package main

import (
	"fmt"

	wcfile "github.com/evanjo03/gowc/internal/file"
	wcflag "github.com/evanjo03/gowc/internal/flag"
	"github.com/evanjo03/gowc/internal/format"
)

func main() {
	flags := wcflag.GetFlags()

	reader, err := wcfile.GetReader()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileCounts := wcfile.GetCounts(reader)
	output := format.GetOutput(flags, fileCounts)

	fmt.Println(output)
}
