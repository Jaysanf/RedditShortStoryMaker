package YoutubeUploader

import (
	"flag"
	"fmt"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
	"strings"
)

type YoutubeUploader interface {
	UploadVideo(filename string)
}

func UploadVideo(filename string, title string) {
	flag.Parse()

	if filename == "" {
		log.Fatalf("You must provide a filename of a video file to upload")
	}

	client := getClient(youtube.YoutubeUploadScope)

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: "#shorts #reddit",
			CategoryId:  "24",
		},
		Status: &youtube.VideoStatus{PrivacyStatus: "public"},
	}
	keywords := "shorts, reddit"
	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(keywords, ",")
	}

	call := service.Videos.Insert([]string{"snippet", "status"}, upload)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", filename, err)
	}

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Unable to post vid", err)
	}
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
}
