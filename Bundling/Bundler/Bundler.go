package Bundler

import (
	"RedditShortStoryMaker/Utils"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"os"
	"time"
)

type Bundler interface {
	Bundle(post *reddit.Post) error
}

func Bundle(post *reddit.Post) error {
	timeStamp := time.Now().Format("01-02-2006_15-04-05")
	workingPath, err := Utils.GetWorkingDirPath()
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(workingPath+dirOutputName+"\\"+timeStamp, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Mkdir(workingPath+dirOutputName+"\\"+timeStamp+"\\mp3", os.ModePerm)
	if err != nil {
		return err
	}

	path := workingPath + dirOutputName + "\\" + timeStamp + "\\"
	err = fractionizePost(path, post)
	if err != nil {
		return err
	}

	//err = mergeMP3FilesIntoOne(path+"/mp3", path+"audio"+mp3File)

	err = getRandomBackgroundVideo(workingPath+dirClipsName, path)
	if err != nil {
		return err
	}

	err = os.RemoveAll(workingPath + dirOutputName + "\\" + timeStamp + "\\mp3")
	if err != nil {
		return err
	}
	return nil
}
