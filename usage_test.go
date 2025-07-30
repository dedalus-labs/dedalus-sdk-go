// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dedalussdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/dedalus-sdk-go"
	"github.com/stainless-sdks/dedalus-sdk-go/internal/testutil"
	"github.com/stainless-sdks/dedalus-sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dedalussdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	response, err := client.Health.Check(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Status)
}
