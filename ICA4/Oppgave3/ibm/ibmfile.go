package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
)

func main() {
	speechToText, speechToTextErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
		IAMApiKey: "iC1tul6yCmfx7zJQnYGB5jFUVDsbKbiL5Cvv0hfXpYK3",
		URL:       "https://gateway-lon.watsonplatform.net/speech-to-text/api",
	})
	if speechToTextErr != nil {
		panic(speechToTextErr)
	}

	// Check for error

	var audioFile io.ReadCloser
	var audioFileErr error
	audioFile, audioFileErr = os.Open("./audio.wav")
	if audioFileErr != nil {
		panic(audioFileErr)
	}
	response, responseErr := speechToText.Recognize(
		&speechtotextv1.RecognizeOptions{
			Audio:       &audioFile,
			ContentType: core.StringPtr(speechtotextv1.RecognizeOptions_ContentType_AudioWav),
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	result := speechToText.GetRecognizeResult(response)
	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))
}
