// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
)

func TestAudioSpeechNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := githubcomdedaluslabsdedalussdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Audio.Speech.New(context.TODO(), githubcomdedaluslabsdedalussdkgo.AudioSpeechNewParams{
		Input:          "Hello, how are you today?",
		Model:          "openai/tts-1",
		Voice:          githubcomdedaluslabsdedalussdkgo.AudioSpeechNewParamsVoiceAlloy,
		Instructions:   githubcomdedaluslabsdedalussdkgo.String("instructions"),
		ResponseFormat: githubcomdedaluslabsdedalussdkgo.AudioSpeechNewParamsResponseFormatMP3,
		Speed:          githubcomdedaluslabsdedalussdkgo.Float(1),
		StreamFormat:   githubcomdedaluslabsdedalussdkgo.AudioSpeechNewParamsStreamFormatSse,
	})
	if err != nil {
		var apierr *githubcomdedaluslabsdedalussdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *githubcomdedaluslabsdedalussdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
