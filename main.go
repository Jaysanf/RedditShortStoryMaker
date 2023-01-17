package main

import (
	"RedditShortStoryMaker/Bundler"
	"RedditShortStoryMaker/ProfanityHandler"
	"RedditShortStoryMaker/RedditHandler"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
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

	ProfanityHandler.RemoveProfanity(&post.Body)
	err = Bundler.Bundle(post)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
	return
}
