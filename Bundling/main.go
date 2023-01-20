package main

import (
	"Bundling/Bundler"
	"Bundling/RedditHandler"
	"Bundling/Utils"
	"fmt"
	"github.com/joho/godotenv"
	"os"
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
	redditHandler := RedditHandler.RedditHandler{}

	err := redditHandler.GetClient(os.Getenv("ID"), os.Getenv("SECRET"),
		os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}
	posts, err := redditHandler.GetTopPosts("tifu", 25, "week")
	if err != nil {
		panic(err)
	}

	usedPostID, err := redditHandler.GetUsedPostID()
	if err != nil {
		panic(err)
	}

	post, err := redditHandler.GetUnusedPost(posts, usedPostID)
	if err != nil {
		panic(err)
	}
	if post == nil {
		panic("No post found")
	}

	//ProfanityHandler.RemoveProfanity(&post.Body)
	err = Bundler.Bundle(post)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
	return
}
