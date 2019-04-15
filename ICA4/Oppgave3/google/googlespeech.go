// Sample speech-quickstart uses the Google Cloud Speech API to transcribe
// audio.
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

<<<<<<< HEAD
//export GOOGLE_APPLICATION_CREDENTIALS="/Users/ingvildtisland/go/src/github.com/IngvildT/IS105-Gruppe13-mappe/ICA-1-3-Repository/ICA4/Oppgave3/google/newjson.json"

=======
// export GOOGLE_APPLICATION_CREDENTIALS="/Users/juliekjellevold/Downloads/api.json"
>>>>>>> f65384e406888ca4d165ccb446a05e36904f7079
func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name of the audio file to transcribe.
<<<<<<< HEAD
	filename := "./test.wav"
=======
	filename := "./audio.wav"
>>>>>>> f65384e406888ca4d165ccb446a05e36904f7079

	// Reads the audio file into memory.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Detects speech in the audio file.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
<<<<<<< HEAD
			SampleRateHertz: 22050,
=======
			SampleRateHertz: 16000,
>>>>>>> f65384e406888ca4d165ccb446a05e36904f7079
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
	}

	// Prints the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
}
