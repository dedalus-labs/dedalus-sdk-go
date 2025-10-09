// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/testutil"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
)

func TestImageGenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Images.Generate(context.TODO(), githubcomdedaluslabsdedalussdkgo.ImageGenerateParams{
		CreateImageRequest: githubcomdedaluslabsdedalussdkgo.CreateImageRequestParam{
			Prompt:            "A white siamese cat",
			Background:        githubcomdedaluslabsdedalussdkgo.CreateImageRequestBackgroundTransparent,
			Model:             githubcomdedaluslabsdedalussdkgo.String("dall-e-3"),
			Moderation:        githubcomdedaluslabsdedalussdkgo.CreateImageRequestModerationAuto,
			N:                 githubcomdedaluslabsdedalussdkgo.Int(1),
			OutputCompression: githubcomdedaluslabsdedalussdkgo.Int(85),
			OutputFormat:      githubcomdedaluslabsdedalussdkgo.CreateImageRequestOutputFormatPng,
			PartialImages:     githubcomdedaluslabsdedalussdkgo.Int(0),
			Quality:           githubcomdedaluslabsdedalussdkgo.CreateImageRequestQualityStandard,
			ResponseFormat:    githubcomdedaluslabsdedalussdkgo.CreateImageRequestResponseFormatURL,
			Size:              githubcomdedaluslabsdedalussdkgo.CreateImageRequestSize1024x1024,
			Stream:            githubcomdedaluslabsdedalussdkgo.Bool(true),
			Style:             githubcomdedaluslabsdedalussdkgo.CreateImageRequestStyleVivid,
			User:              githubcomdedaluslabsdedalussdkgo.String("user"),
		},
	})
	if err != nil {
		var apierr *githubcomdedaluslabsdedalussdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
