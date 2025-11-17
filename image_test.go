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

func TestImageNewVariationWithOptionalParams(t *testing.T) {
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
	_, err := client.Images.NewVariation(context.TODO(), githubcomdedaluslabsdedalussdkgo.ImageNewVariationParams{
		Image:          githubcomdedaluslabsdedalussdkgo.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		Model:          githubcomdedaluslabsdedalussdkgo.F("model"),
		N:              githubcomdedaluslabsdedalussdkgo.F(int64(0)),
		ResponseFormat: githubcomdedaluslabsdedalussdkgo.F("response_format"),
		Size:           githubcomdedaluslabsdedalussdkgo.F("size"),
		User:           githubcomdedaluslabsdedalussdkgo.F("user"),
	})
	if err != nil {
		var apierr *githubcomdedaluslabsdedalussdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestImageEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Images.Edit(context.TODO(), githubcomdedaluslabsdedalussdkgo.ImageEditParams{
		Image:          githubcomdedaluslabsdedalussdkgo.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		Prompt:         githubcomdedaluslabsdedalussdkgo.F("prompt"),
		Mask:           githubcomdedaluslabsdedalussdkgo.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		Model:          githubcomdedaluslabsdedalussdkgo.F("model"),
		N:              githubcomdedaluslabsdedalussdkgo.F(int64(0)),
		ResponseFormat: githubcomdedaluslabsdedalussdkgo.F("response_format"),
		Size:           githubcomdedaluslabsdedalussdkgo.F("size"),
		User:           githubcomdedaluslabsdedalussdkgo.F("user"),
	})
	if err != nil {
		var apierr *githubcomdedaluslabsdedalussdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

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
			Prompt:            githubcomdedaluslabsdedalussdkgo.F("A white siamese cat"),
			Background:        githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.CreateImageRequestBackgroundTransparent),
			Model:             githubcomdedaluslabsdedalussdkgo.F("openai/dall-e-3"),
			Moderation:        githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.CreateImageRequestModerationAuto),
			N:                 githubcomdedaluslabsdedalussdkgo.F(int64(1)),
			OutputCompression: githubcomdedaluslabsdedalussdkgo.F(int64(85)),
			OutputFormat:      githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.CreateImageRequestOutputFormatPng),
			PartialImages:     githubcomdedaluslabsdedalussdkgo.F(int64(0)),
			Quality:           githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.CreateImageRequestQualityStandard),
			ResponseFormat:    githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.CreateImageRequestResponseFormatURL),
			Size:              githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.CreateImageRequestSize1024x1024),
			Stream:            githubcomdedaluslabsdedalussdkgo.F(true),
			Style:             githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.CreateImageRequestStyleVivid),
			User:              githubcomdedaluslabsdedalussdkgo.F("user"),
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
