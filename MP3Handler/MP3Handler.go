package MP3Handler

import (
	"github.com/aws/aws-sdk-go/aws"
	session2 "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"io"
	"os"
)

type PollyService interface {
	Synthesize(text string, filename string) error
}

type pollyConfig struct {
	voice string
}

func NewPollyService(voice personsVoice) PollyService {
	return &pollyConfig{
		voice: string(voice),
	}
}

func createPollyClient() *polly.Polly {
	// Create AWS session
	session := session2.Must(session2.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	}))
	return polly.New(session)
}

func (config *pollyConfig) Synthesize(text string, filename string) error {
	pollyClient := createPollyClient()

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String(text),
		VoiceId:      aws.String(config.voice),
	}

	output, err := pollyClient.SynthesizeSpeech(input)
	if err != nil {
		return err
	}

	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer outFile.Close() // close files when func end

	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		return err
	}

	return nil
}
