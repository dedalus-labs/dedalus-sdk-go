// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo_test

import (
	"context"
	"os"
	"testing"

	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/testutil"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/shared"
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
		Model: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsModelUnion](shared.UnionString("openai/gpt-5-nano")),
		Messages: githubcomdedaluslabsdedalussdkgo.F[githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMessagesUnion](githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParamsMessagesArray([]map[string]interface{}{{
			"role":    "user",
			"content": "Hello, how are you today?",
		}})),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", completion.ID)
}
