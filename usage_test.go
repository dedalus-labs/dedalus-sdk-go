// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo_test

import (
	"context"
	"os"
	"testing"

	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/testutil"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
)

func TestUsage(t *testing.T) {
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
	completion, err := client.Chat.Completions.New(context.TODO(), githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParams{
		Model: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsModelUnion{
			OfModelID: githubcomdedaluslabsdedalussdkgo.String("openai/gpt-4o-mini"),
		},
		Messages: githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMessagesUnion{
			OfMapOfAnyMap: []map[string]any{{
				"role":    "user",
				"content": "Hello, how are you today?",
			}},
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", completion.ID)
}
