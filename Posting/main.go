package main

import (
	"Posting/Utils"
	"Posting/YoutubeUploader"
	"github.com/joho/godotenv"
)

func init() {
	workingPath, err := Utils.GetWorkingDirPath()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(workingPath + "\\.env")
	if err != nil {
		panic(err)
	}
}

func main() {
	workingPath, err := Utils.GetWorkingDirPath()
	if err != nil {
		panic(err)
	}

	bundlePath, err := Utils.GetLatestBundle(workingPath + "\\Shorts")
	if err != nil {
		panic(err)
	}

	post, err := Utils.GetPost(bundlePath)
	if err != nil {
		panic(err)
	}

	YoutubeUploader.UploadVideo(bundlePath+"\\final.mp4", post.Title[0:Utils.Min(100, len(post.Title))])
}
