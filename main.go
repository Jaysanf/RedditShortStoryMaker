package main

import (
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
		fmt.Printf("%v \n", err)
		return
	}
	redditHandler.GetTopPosts("golang")
	fmt.Println("Done")
}
