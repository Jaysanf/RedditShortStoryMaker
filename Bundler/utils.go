package Bundler

import (
	"io"
	"os"
	"strings"
)

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func divideText(text string, x int) []string {
	words := strings.Fields(text)
	var dividedText []string
	for i := 0; i < len(words); i += x {
		j := i + x
		if j > len(words) {
			j = len(words)
		}
		dividedText = append(dividedText, strings.Join(words[i:j], " "))
	}
	return dividedText
}
