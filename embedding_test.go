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
	"github.com/dedalus-labs/dedalus-sdk-go/shared"
)

func TestEmbeddingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Embeddings.New(context.TODO(), githubcomdedaluslabsdedalussdkgo.EmbeddingNewParams{
		CreateEmbeddingRequest: githubcomdedaluslabsdedalussdkgo.CreateEmbeddingRequestParam{
			Input:          githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.CreateEmbeddingRequestInputUnionParam](shared.UnionString("string")),
			Model:          githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.CreateEmbeddingRequestModelTextEmbeddingAda002),
			Dimensions:     githubcomdedaluslabsdedalussdkgo.F(int64(1)),
			EncodingFormat: githubcomdedaluslabsdedalussdkgo.F(githubcomdedaluslabsdedalussdkgo.CreateEmbeddingRequestEncodingFormatFloat),
			User:           githubcomdedaluslabsdedalussdkgo.F("user"),
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
