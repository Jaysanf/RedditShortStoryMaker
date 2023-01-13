package Bundler

import (
	"RedditShortStoryMaker/MP3Handler"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"math"
	"os"
	"time"
)

type Bundler interface {
	Bundle(post *reddit.Post) error
}

func Bundle(post *reddit.Post) error {
	timeStamp := time.Now().Format("01-02-2006_15-04-05")
	err := os.Mkdir("Shorts/"+timeStamp, os.ModePerm)
	if err != nil {
		return err
	}
	path := dirOutputName + "/" + timeStamp + "/"

	mp3Handler := MP3Handler.NewPollyService(MP3Handler.Matthew)
	speech := post.Title + post.Body
	err = mp3Handler.Synthesize(speech[:int(math.Min(2000, float64(len(post.Body))))], path+mp3Name)
	if err != nil {
		return err
	}

	f, err := os.Create(path + textName)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.WriteString(post.Body)
	if err != nil {
		return err
	}
	return nil
}
