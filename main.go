package main

import (
	"RedditShortStoryMaker/MP3Handler"
	"RedditShortStoryMaker/RedditHandler"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math"
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
	mp3Handler := MP3Handler.NewPollyService(MP3Handler.Matthew)
	fmt.Printf("%v", len(post.Title+" \n"+post.Body))
	fmt.Printf("%v", (post.Title + " \n" + post.Body))

	err = mp3Handler.Synthesize(post.Body[:int(math.Min(2000, float64(len(post.Body))))], "test.mp3")
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")

}
