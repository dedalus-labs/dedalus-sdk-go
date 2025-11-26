// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/testutil"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
)

func TestAudioTranscriptionNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomdedaluslabsdedalussdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Audio.Transcriptions.New(context.TODO(), githubcomdedaluslabsdedalussdkgo.AudioTranscriptionNewParams{
		File:           githubcomdedaluslabsdedalussdkgo.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		Model:          githubcomdedaluslabsdedalussdkgo.F("model"),
		Language:       githubcomdedaluslabsdedalussdkgo.F("language"),
		Prompt:         githubcomdedaluslabsdedalussdkgo.F("prompt"),
		ResponseFormat: githubcomdedaluslabsdedalussdkgo.F("response_format"),
		Temperature:    githubcomdedaluslabsdedalussdkgo.F(0.000000),
	})
	if err != nil {
		var apierr *githubcomdedaluslabsdedalussdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
