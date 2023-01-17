package main

import (
	"RedditShortStoryMaker/Bundler"
	"RedditShortStoryMaker/ProfanityHandler"
	"RedditShortStoryMaker/RedditHandler"
	"fmt"
	"github.com/cznic/ql"
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
	// Open a connection to a SQLite3 database

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
		panic("No post found")
	}
	InitDb()

	return
	ProfanityHandler.RemoveProfanity(&post.Body)
	err = Bundler.Bundle(post)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
	return
}

func InitDb() {
	// Create a new table
	db, err := ql.OpenFile("mydb.db", &ql.Options{CanCreate: true})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, _, err = db.Run(ql.NewRWCtx(), `
										BEGIN TRANSACTION;
											CREATE TABLE  IF NOT EXISTS department (
												DepartmentID   int,
												DepartmentName string,
											);
											CREATE TABLE employee (
												LastName	string,
												DepartmentID	int,
											);
										COMMIT;
									`, nil)
	if err != nil {
		panic(err)
	}

	// Select data from the table
	if err != nil {
		panic(err)
	}
	fmt.Println(rs)
}
