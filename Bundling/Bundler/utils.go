package Bundler

import (
	MP3Handler2 "RedditShortStoryMaker/MP3Handler"
	"fmt"
	"github.com/hyacinthus/mp3join"
	"github.com/tcolgate/mp3"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"golang.org/x/exp/rand"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Create multiple mp3 and an SRT file for the reddit post
func fractionizePost(path string, post *reddit.Post) error {
	bodyFractionized := divideText(post.Body, numberOfWordsPerSplit)
	bodyFractionized = append([]string{post.Title}, bodyFractionized...) // Adding the Title

	// Create a new SRT file to write the subtitles to
	srt, err := os.Create(path + "subtitles.srt")
	if err != nil {
		return err
	}
	defer srt.Close()
	// Init var
	subtitleNum := 1
	startTime := time.Duration(0)
	endTime := time.Duration(0)

	mp3Handler := MP3Handler2.NewPollyService(MP3Handler2.Matthew)
	for i, chunkOfWords := range bodyFractionized {
		fileNameMP3 := path + "mp3/" + strconv.Itoa(i)

		err := mp3Handler.Synthesize(chunkOfWords, fileNameMP3+mp3File)
		if err != nil {
			return err
		}
		duration, err := getDurationOfMp3File(fileNameMP3 + mp3File)
		if err != nil {
			return err
		}
		endTime += time.Duration(duration * float64(time.Second))

		startTimeStr := fmtDuration(startTime)

		endTimeStr := fmtDuration(endTime)

		// Write to srt file
		_, err = fmt.Fprintln(srt, subtitleNum)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(srt, "%s --> %s\n", startTimeStr, endTimeStr)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(srt, chunkOfWords)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(srt, "") // New line at the end
		if err != nil {
			return err
		}
		// Increment
		startTime = endTime
		subtitleNum++
	}
	text := post.Title + ".\n" + post.Body
	err = mp3Handler.Synthesize(text[:Min(len(text), 3000)], path+"audio"+mp3File)
	if err != nil {
		return err
	}

	return nil
}

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
	defer r.Close()

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

func mergeMP3FilesIntoOne(mp3Dir string, fileName string) error {
	joiner := mp3join.New()
	files, err := os.ReadDir(mp3Dir)
	if err != nil {
		return err
	}

	sort.Slice(files, func(i, j int) bool {
		x, _ := strconv.Atoi(files[i].Name()[:len(files[i].Name())-len(filepath.Ext(files[i].Name()))])
		y, _ := strconv.Atoi(files[j].Name()[:len(files[j].Name())-len(filepath.Ext(files[j].Name()))])
		return x < y
	})
	// readers is the input mp3 files
	for _, file := range files {
		mp3Byte, err := os.Open(mp3Dir + "/" + file.Name())
		if err != nil {
			return err
		}

		err = joiner.Append(mp3Byte)
		if err != nil {
			return err
		}
		err = mp3Byte.Close()
		if err != nil {
			return err
		}
	}

	dest := joiner.Reader()
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = dest.WriteTo(f)
	if err != nil {
		return err
	}

	return nil
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
