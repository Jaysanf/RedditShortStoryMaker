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

	post := redditHandler.GetUnusedPost(posts, []string{})
	if post == nil {
		panic(err)
	}

	ProfanityHandler.RemoveProfanity(&post.Body)
	Bundler.Bundle(post)
	fmt.Println("Done")
	return
}
