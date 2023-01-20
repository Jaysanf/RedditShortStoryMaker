package Utils

import (
	"os"
	"path/filepath"
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
