package Utils

import (
	"golang.org/x/exp/rand"
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
	if workingPath[len(workingPath)-len("Bundling"):] == "Bundling" {
		return workingDir, nil
	}

	return workingPath, nil
}

var subReddits = []string{"tifu", "amitheasshole"}

func GetRandomSubreddit() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	return subReddits[rand.Intn(len(subReddits))]
}
