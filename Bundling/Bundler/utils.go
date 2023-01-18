package Bundler

import (
	"fmt"
	"github.com/tcolgate/mp3"
	"golang.org/x/exp/rand"
	"io"
	"os"
	"strings"
	"time"
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

func fmtDuration(d time.Duration) string {
	h := d / time.Hour
	d -= h * time.Hour

	m := d / time.Minute
	d -= m * time.Minute

	s := d / time.Second
	d -= s * time.Second

	ms := d / time.Millisecond
	d -= ms * time.Millisecond
	return fmt.Sprintf("%02d:%02d:%02d,%03d", h, m, s, ms)
}

func getRandomBackgroundVideo(videoDirName, copyPlaceDirName string) error {
	files, err := os.ReadDir(videoDirName)
	if err != nil {
		return err
	}
	rand.Seed(uint64(time.Now().UnixNano()))
	randomVideo := files[(rand.Intn(len(files)-1) + 1)] // Get rand video from 1 to n -1, exclude .gitkeep
	err = copyFileContents(videoDirName+"/"+randomVideo.Name(), copyPlaceDirName+"/video"+mp4File)
	if err != nil {
		return err
	}

	return nil
}

func getDurationOfMp3File(mp3File string) (float64, error) {
	t := 0.0

	r, err := os.Open(mp3File)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	d := mp3.NewDecoder(r)
	var f mp3.Frame
	skipped := 0

	for {
		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			return -1, err
		}
		t = t + f.Duration().Seconds()
	}

	return t, nil
}
