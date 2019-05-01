package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
)

type SpeechtoText struct {
	Transcript string  `json:"transcript"`
	Results []struct {
		Final        bool `json:"final"`
		Alternatives []struct {
			Transcript string  `json:"transcript"`
			Confidence float64 `json:"confidence"`
		} `json:"alternatives"`
	} `json:"results"`
	ResultIndex int `json:"result_index"`
}

func main() {
	speechToText, speechToTextErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
		IAMApiKey: "tHAIyaaB5VqOFBRZE8asIEF37i5nTtUkon9a1T3x8Rc4",
		URL:       "https://gateway-lon.watsonplatform.net/speech-to-text/api",
	})
	if speechToTextErr != nil {
		panic(speechToTextErr)
	}

	// Check for error

	var audioFile io.ReadCloser
	var audioFileErr error
	audioFile, audioFileErr = os.Open("./audio-file.flac")
	if audioFileErr != nil {
		panic(audioFileErr)
	}
	response, responseErr := speechToText.Recognize(
		&speechtotextv1.RecognizeOptions{
			Audio:             &audioFile,
			ContentType:       core.StringPtr(speechtotextv1.RecognizeOptions_ContentType_AudioFlac),
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}

	result := speechToText.GetRecognizeResult(response)


	var text SpeechtoText
	err := json.Unmarshal(result, &text)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", text.Transcript)


	

	//text ,_ := json.MarshalIndent(result, "", "  ")

	//fmt.Println(string(text))
