package Bundler

import (
	MP3Handler2 "RedditShortStoryMaker/MP3Handler"
	"fmt"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"os"
	"strconv"
	"time"
)

type Bundler interface {
	Bundle(post *reddit.Post) error
}

func Bundle(post *reddit.Post) error {
	timeStamp := time.Now().Format("01-02-2006_15-04-05")
	err := os.Mkdir(dirOutputName+"/"+timeStamp, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Mkdir(dirOutputName+"/"+timeStamp+"/mp3", os.ModePerm)
	if err != nil {
		return err
	}

	path := dirOutputName + "/" + timeStamp + "/"
	err = fractionizePost(path, post)
	if err != nil {
		return err
	}

	err = getRandomBackgroundVideo(dirClipsName, path)
	if err != nil {
		return err
	}

	return nil
}

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

	return nil
}
