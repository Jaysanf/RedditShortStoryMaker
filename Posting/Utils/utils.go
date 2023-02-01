package Utils

import (
	"encoding/json"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"os"
	"path/filepath"
	"time"
)

type Utils interface {
	GetWorkingDirPath() (string, error)
}

func GetWorkingDirPath() (string, error) {
	workingPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	workingDir := filepath.Dir(workingPath)

	// Means we are running from the dir bundling
	if workingPath[len(workingPath)-len("Posting"):] == "Posting" {
		return workingDir, nil
	}

	return workingPath, nil
}

func GetPost(path string) (*reddit.Post, error) {
	f, err := os.ReadFile(path + "\\post.txt")
	if err != nil {
		return nil, err
	}
	post := reddit.Post{}

	err = json.Unmarshal(f, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetLatestBundle(path string) (string, error) {
	d, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var latestTimestamp time.Time
	var latestTimestampDir string
	for _, file := range d {
		if !file.IsDir() {
			continue
		}

		timestampStr := file.Name()
		timestamp, err := time.Parse("01-02-2006_15-04-05", timestampStr)
		if err != nil {
			continue
		}

		if timestamp.After(latestTimestamp) {
			latestTimestamp = timestamp
			latestTimestampDir = file.Name()
		}

	}

	return path + "\\" + latestTimestampDir, nil
}
