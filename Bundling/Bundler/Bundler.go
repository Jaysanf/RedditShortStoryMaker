package Bundler

import (
	MP3Handler2 "RedditShortStoryMaker/MP3Handler"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"golang.org/x/exp/rand"
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

// Create multiple mp3 and txt files for the reddit post
func fractionizePost(path string, post *reddit.Post) error {
	bodyFractionized := divideText(post.Body, numberOfWordsPerSplit)

	bodyFractionized = append([]string{post.Title}, bodyFractionized...) // Adding the Title
	mp3Handler := MP3Handler2.NewPollyService(MP3Handler2.Matthew)
	for i, chunkOfWords := range bodyFractionized {
		fileName := path + strconv.Itoa(i)
		err := mp3Handler.Synthesize(chunkOfWords, fileName+mp3File)
		if err != nil {
			return err
		}

		f, err := os.Create(fileName + txtFile)
		if err != nil {
			return err
		}

		defer f.Close()
		_, err = f.WriteString(chunkOfWords)
		if err != nil {
			return err
		}
	}

	return nil
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
