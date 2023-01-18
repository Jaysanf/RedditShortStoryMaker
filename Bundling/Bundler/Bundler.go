package Bundler

import (
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"os"
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

	//err = mergeMP3FilesIntoOne(path+"/mp3", path+"audio"+mp3File)

	err = getRandomBackgroundVideo(dirClipsName, path)
	if err != nil {
		return err
	}

	err = os.RemoveAll(dirOutputName + "/" + timeStamp + "/mp3")
	if err != nil {
		return err
	}
	return nil
}
